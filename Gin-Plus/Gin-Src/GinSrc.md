# Gin框架源码解析

## 路由源码解析

gin 框架使用的是定制版的httprouter，其路由的原理是大量使用公共前缀的树结构。具有公共前缀的节点也共享一个公共父节点\

> Radix Tree

基数树 是一种更节省空间的前缀树，对于基数树的每一个节点，如果该节点是唯一的子树的话，就和父节点合并。

~~~ go
func main() {

	r := gin.Default()

	r.GET("/", func1)
	r.GET("/search/", func2)
	r.GET("/support/", func3)
	r.GET("/blog/", func4)
	r.GET("/blog/:post/", func5)
	r.GET("/about-us/", func6)
	r.GET("/about-us/team/", func7)

	r.Run(":8080")

}
~~~


~~~ bash
优先级 路径             处理函数
9       /               *<1>
3       |-s             nil
2       | |-earch/      *<2>
1       | |-upport/     *<3>
2       |blog/          *<4>
1       |    \:post     nil
1       |          \/   *<5> 
2       |-about-us/     *<6>
1       |         \team *<7>
~~~

上边最右边那一列每一个*<num>表示处理函数的内存地址 从根节点遍历到叶子节点我们就能得到完整的路由表

由于URL路径具有层次结构，并且只使用有限的一组字符，所以可能有许多常见的前缀，这使得我们可以很容易地将路由简化为更小的问题
此外，路由器为每一个请求方法管理一颗单独的树。一方面，它在每一个节点中都保存一个 method -> handle map更加节省空间，它还使我们甚至可以在开始在前缀树中查找之前大大减少路由问题。

为了获得更好的可伸缩性，每一级的子节点都按`优先级`排序，其中优先级就是子节点中注册的句柄的数量
- 首先优先匹配被大多数路由路径包含的节点，遮阳可以让尽可能多的路由快速被定位
- 类似于成本补偿，最长的路径可以被优先匹配，补偿体现在最长的路径需要更长的时间来定位，如果最长路径的节点能被优先匹配，那么路由匹配所花的时间不一定比短路径的路由时间长

### 解析
> 初始化

~~~ go
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}
~~~

~~~ go
type Engine struct {
	RouterGroup

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	HandleMethodNotAllowed bool
	ForwardedByClientIP    bool

	// #726 #755 If enabled, it will thrust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
	AppEngine bool

	// If enabled, the url.RawPath will be used to find parameters.
	UseRawPath bool

	// If true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
	UnescapePathValues bool

	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
	MaxMultipartMemory int64

	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
	// See the PR #1817 and issue #1644
	RemoveExtraSlash bool

	delims           render.Delims
	secureJsonPrefix string
	HTMLRender       render.HTMLRender
	FuncMap          template.FuncMap
	allNoRoute       HandlersChain
	allNoMethod      HandlersChain
	noRoute          HandlersChain
	noMethod         HandlersChain
	pool             sync.Pool
	trees            methodTrees
}
~~~


// 提高性能： 通过对象池 减少每一次临时申请对象 GC 的资源占用
~~~ go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)//sync poll 对象池 减少GC 减少内存申请 然后断言
	c.writermem.reset(w) //取出后再初始化
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}
~~~

- 处理请求函数
~~~ go
func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false
	if engine.UseRawPath && len(c.Request.URL.RawPath) > 0 {
		rPath = c.Request.URL.RawPath
		unescape = engine.UnescapePathValues
	}

	if engine.RemoveExtraSlash {
		rPath = cleanPath(rPath)
	}

	// Find root of the tree for the given HTTP method
	t := engine.trees
    // for i := 0;i < len(t);i++  计算一次len(t)即可
	for i, tl := 0, len(t); i < tl; i++ {
        // 最小代码原则 将停止循环的条件放到最前面
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// Find route in tree
		value := root.getValue(rPath, c.Params, unescape)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.Params = value.params
			c.fullPath = value.fullPath
			c.Next()
			c.writermem.WriteHeaderNow()
			return
		}
		if httpMethod != "CONNECT" && rPath != "/" {
			if value.tsr && engine.RedirectTrailingSlash {
				redirectTrailingSlash(c)
				return
			}
			if engine.RedirectFixedPath && redirectFixedPath(c, root, engine.RedirectFixedPath) {
				return
			}
		}
		break
	}

	if engine.HandleMethodNotAllowed {
		for _, tree := range engine.trees {
			if tree.method == httpMethod {
				continue
			}
			if value := tree.root.getValue(rPath, nil, unescape); value.handlers != nil {
				c.handlers = engine.allNoMethod
				serveError(c, http.StatusMethodNotAllowed, default405Body)
				return
			}
		}
	}
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}
~~~

**确保结构体实现了指定的接口 因为如果没有实现的话 编译阶段会报错**
~~~ go

type Engine struct {...}

var _ IRouter = &Engine{}// 匿名变量 确保结构体实现了指定的接口 因为如果没有实现的话 编译阶段会报错
~~~

- engine.trees

~~~ go
type Engine struct {
	trees            methodTrees
}

~~~

~~~ go

type methodTrees []methodTree

type methodTree struct {
	method string
	root   *node
}


// http1.1 只有几种请求方法 没必要节省时间去使用更耗内存的hash
func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}
~~~

- 方法初始化

~~~ go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
~~~

- 引擎初始化
~~~ go
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		AppEngine:              defaultAppEngine,
		UseRawPath:             false,
		RemoveExtraSlash:       false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJsonPrefix:       "while(1);",
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}
~~~

`trees: make(methodTrees, 0, 9)`

- 初始化结束后 返回结果
~~~ go
type node struct {
    // 节点路径 比如上面的s earch upport
	path      string
    // 与children字段对应 保存的是分裂的分支的第一个字符
    // 例如search 和support 那么s 节点的indices对应的eu
    // 代表有两个分支 分支的首字母分别是e 和 u
	indices   string
    // 子节点
	children  []*node
    // 处理函数链条(slice)
	handlers  HandlersChain
    // 优先级 子节点 子子节点 等注册的handler数量
	priority  uint32
    // 节点类型 包括static root param catchAll
    // static 静态节点 比如上面的 s earch 等节点 
    // root 树的根节点
    // catchAdd 有*匹配的节点
    // param 参数节点
	nType     nodeType
    /// 路径上最大参数的个数
	maxParams uint8
    // 节点上是否是参数节点 比如 :post
	wildChild bool
    // 完整路径
	fullPath  string
}
~~~

> 注册路由

注册路由的逻辑主要有 `addRoute` 函数和`insertChild` 方法

~~~ go
r.GET("/", func1)
~~~

~~~ go
// GET is a shortcut for router.Handle("GET", path, handle).
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handlers)
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
    // 处理函数拼接过程
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

~~~

~~~ go
// addRoute 将具有给定句柄的节点添加到路径中
// 非并发安全 => 在goroutine中使用时要注意
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path //记录fullpath
	n.priority++ //将当前优先级提升，注册了一个新的处理函数
	numParams := countParams(path) //数一下参数个数

	// 空树的话就直接插入当前节点
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk: // label 表示比较复杂的嵌套
	for {
		// 更新当前节点的最大参数个数
		if numParams > n.maxParams {
			n.maxParams = numParams
		}

		//  寻找最长的通用前缀
		// 这也意味着公共前缀不能包含 : * 
		// 因为现有键不能包含这些字符
		i := longestCommonPrefix(path, n.path)

		// 分裂边缘 此处分裂的是当前树的节点
        // 例如一开始path 是 search 新加入support s是它们通用的最长前缀部分
        // 那么会将s拿出来作为parent节点 增加earch 和 upport作为child节点
		if i < len(n.path) {
			child := node{
				path:      n.path[i:], // 公共前缀后的部分作为子节点
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1, // 子节点优先级-1
				fullPath:  n.fullPath,
			}

			// Update maxParams (max of all children)
			for _, v := range child.children {
				if v.maxParams > child.maxParams {
					child.maxParams = v.maxParams
				}
			}

			n.children = []*node{&child}
			// 保存第一个字符 e u 
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

        // 第二分支
		// 将新来的节点插入新的oarent节点作为子节点
		if i < len(path) {
			path = path[i:]

			if n.wildChild { // 如果是参数节点
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++

				// Update maxParams of the child node
				if numParams > n.maxParams {
					n.maxParams = numParams
				}
				numParams--

				// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
					// 检查更长的通配符 例如 :name :names
					if len(n.path) >= len(path) || path[len(n.path)] == '/' {
						continue walk
					}
				}

				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(path, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}
            
            // 取path首字母 用来与indices做比较
			c := path[0]

			// 处理参数后加斜线的情况
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			// 检查path下一个子节点的子节点是否存在
            // 比如s的子节点现在是earch upport  indices为eu
            // 比如新加一个路由为super 那么就是和 upper 有匹配的部分u 将继续分裂现在的 upport节点
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			// 否则就插入
			if c != ':' && c != '*' {
				// []byte for proper unicode char conversion, see #65
				n.indices += string([]byte{c})
				child := &node{
					maxParams: numParams,
					fullPath:  fullPath,
				}
				n.children = append(n.children, child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			}
			n.insertChild(numParams, path, fullPath, handlers)
			return
		}

		// 已经注册过的节点
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		return
	}
}
~~~

> 中间件

gin 的中间件设计的十分巧妙 ,让我们从`Logger` 和 `Recovery`中间件开始

~~~ go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
~~~

~~~ go

func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}
~~~

~~~ go
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}
~~~





## 

## 

