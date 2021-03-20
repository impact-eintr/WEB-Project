## 模板

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
{{ $v1 := 100 }}
{{ if $v1 }}
{{ $v1 }}
{{ else }}
test
{{ end }}

{{ range $idx,$val := . }}
    <p>{{ $idx }} - {{ $val.Name }}</p>
{{ end }}

</body>
</html>
~~~

~~~ go
import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type User struct {
	Name   string
	gender string
	Age    int
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err :%s\n", err)
		return
	}

	u1 := User{
		Name:   "eintr",
		gender: "男",
		Age:    23,
	}

	u2 := User{
		Name:   "szc",
		gender: "男",
		Age:    22,
	}

	//users := map[string]User{
	//	"u1": u1,
	//	"u2": u2,
	//}
	users := []User{u1, u2}

	err = t.Execute(w, users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err :%s\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err)
		return
	}
}
~~~

### 自定义模板函数

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>

{{ fuck . }}

</body>
</html>
~~~

~~~ go
func f(w http.ResponseWriter, r *http.Request) {
	t := template.New("f.tmpl")

	tfuck := func(name string) (string, error) {
		return name + "NB!", nil
	}

	t.Funcs(template.FuncMap{
		"fuck": tfuck,
	})

	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}

	name := "eintr"
	t.Execute(w, name)
}
~~~

### 模板嵌套
~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>测试嵌套模板</title>
</head>
<body>

<h1>测试嵌套模板<h2>

<hr>
{{ template "ul.tmpl" }}
<hr>

<hr>
{{ template "ol.tmpl" }}
<hr>

你好，{{ . }}

</body>
</html>

{{ define "ol.tmpl" }}
<ol>
    <li>吃饭饭</li>
    <li>睡觉觉</li>
    <li>打豆豆</li>
</ol>
{{ end }}
~~~

~~~ html
<ul>
    <li>吃饭饭</li>
    <li>睡觉觉</li>
    <li>打豆豆</li>
</ul>

~~~

~~~ go
func t(w http.ResponseWriter, r *http.Request) {
	t := template.New("t.tmpl")

	_, err := t.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}

	name := "eintr"
	t.Execute(w, name)
}
~~~

### 模板继承

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板继承</title>
    <style>
    * {
    margin: 0;
    }
    .nav {
        width: 100%;
        height: 50px;
        position: fixed;
        top: 0;
        background-color: burlywood;
    }
    
    .main {
        margin-top: 50px;
    }
    
    .menu {
       width: 15%;
       height: 100%;
       position: fixed;
       left: 0;
       background-color: #6ad7e5;
    }
    
    .center {
        text-align: center;
    }

</style>
</head>
<body>
    <div class="nav"><div>
    <div class="main">
        <div class="menu"></div>
        <div class="content center">
            {{ block "content" . }}{{ end }}
        </div>
    </div>

</body>
</html>

~~~

~~~ html
{{ template "base.tmpl" }}
{{ define "content" }}
    <h1>这是家目录</h1>
    <h2>hello {{ . }}</h2>
{{ end }}

~~~


~~~ html
{{ template "base.tmpl" }}
{{ define "content" }}
    <h1>这是首页</h1>
    <h2>hello {{ . }}</h2>
{{ end }}

~~~

~~~ go
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse err :%s\n", err)
		return
	}
	msg := "eintr"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse err :%s\n", err)
		return
	}
	msg := "eintr"
	t.Execute(w, msg)

}

func main() {
	http.HandleFunc("/", t)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err)
		return
	}
}

~~~

> 避免XSS攻击

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板继承</title>
    <style>
    * {
    margin: 0;
    }
    .nav {
        width: 100%;
        height: 50px;
        position: fixed;
        top: 0;
        background-color: burlywood;
    }
    
    .main {
        margin-top: 50px;
    }
    
    .menu {
       width: 15%;
       height: 100%;
       position: fixed;
       left: 0;
       background-color: #6ad7e5;
    }
    
    .center {
        text-align: center;
    }

</style>
</head>
<body>
    <div class="nav"><div>
    <div class="main">
        <div class="menu"></div>
        <div class="content center">
            {{ .str1 }}
            {{ .str2 | safe }}
        </div>
    </div>

</body>
</html>
~~~

~~~ go
func Xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("Xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("Xss.tmpl")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	str1 := "<script>alert('123');</script>"
	str2 := "<a href='http://127.0.0.1:8081/xss'>回环测试</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
~~~

### Gin自动绑定
~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <form action="/user" method="post">
        <div>
            用户名：
            <input type="text" name="username">
        </div>
        <div>
            密码：
            <input type="text" name="password">
        </div>
        <div>
            <input type="submit" value="submit">
        </div>
    </form>
</body>
</html>
~~~

~~~ go

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":8081")
}
~~~

## Gin处理Json

~~~ go
func main() {
	r := gin.Default()
	type msg struct {
		Name string `json:"name"`
		Word string `json:"word"`
	}

	r.GET("/json", func(c *gin.Context) {
		data := msg{
			Name: "eintr",
			Word: "wdnm!",
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":8081")
}
~~~

## Gin处理Query

~~~ html

~~~

~~~ go
func main() {
	r := gin.Default()

	// /?key1=val1&key2=val2&key3=val3
	r.GET("/", func(c *gin.Context) {
		//name := c.Query("query")
		//name := c.DefaultQuery("query","none")
		name, ok := c.GetQuery("query")
		if !ok {
			log.Println("err")
			name = "none"
		}
		age, ok := c.GetQuery("age")
		if !ok {
			log.Println("err")
			name = "none"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8081")
}
~~~

## Gin处理Form

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>
    <h1>hello ! {{ .name }}</h1>
    
</body>
</html>
~~~

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <form action="/login" method="post">

        <div>
            <label for="username">username:</label>
            <input type="text" name="username" id="username">
        </div>

        <div>
            <label for="">password:</label>
            <input type="password" name="password" id="password">
        </div>

        <div>
            <input type="submit" value="Login">
        </div>
    </form>
</body>
</html>
~~~

~~~ go

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":     username,
			"password": password,
		})
	})

	r.Run(":8081")
}
~~~


## Gin解析URL

~~~ go
func main() {
	r := gin.Default()

	r.GET("/user/:id/:age", func(c *gin.Context) {
		name := c.Param("id")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8081")
}
~~~

##  Gin重定向
~~~ go
func main() {

	r := gin.Default()

	//请求重定向
	r.GET("index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	//请求转发
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //将请求的URL修改
		r.HandleContext(c)        //继续之后的操作
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":8081")
}
~~~

## Gin上传文件

~~~ html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="f1">
        <input type="submit" value="upload">
    </form>
</body>
</html>
~~~

~~~ go
func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			dst := fmt.Sprintf("./%s", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})

	r.Run(":8081")
}
~~~

## Gin路由处理

~~~ go
func main() {
	r := gin.Default()

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"method": "Delete",
			})
		}
	})

	// 路由组
	// 将共用的前缀提取出来 创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video index",
			})
		})
		videoGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video home",
			})
		})
	}

	// 非法路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "-_-# 页面找不到呀",
		})
	})
	r.Run(":8081")

}
~~~

## 中间件

~~~ go
func indexHandler(c *gin.Context) {
	name, _ := c.Get("name") //从上下文中取值（跨中间件取值）
	c.JSON(http.StatusOK, gin.H{
		"msg":  "This is the index page",
		"name": name,
	})
}

func homeHandler(c *gin.Context) {
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "This is the home page",
		"name": name,
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 is working...")
	start := time.Now()
	c.Next() //调用后续的处理函数 即indexHandler...
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 exited...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 is working...")
	c.Next() //调用后续的处理函数 即indexHandler...
	fmt.Println("m2 exited...")
}

func m3(c *gin.Context) {
	c.Set("name", "eintr") //在上下文中设置值
}

func main() {
	r := gin.Default() //默认使用了Logger() 和 Recover()中间件
	//r := gin.New() //不包含任何中间件
	//当中间件或者handler中启动新的goroutine时不能使用原始上下文(c *gin.Context)必须使用其只读副本(c.Copy())
	userGroup := r.Group("/user")

	//为路由组设置中间件
	userGroup.Use(m1, m2, m3)

	userGroup.GET("/index", indexHandler)
	userGroup.GET("/home", homeHandler)

	r.Run(":8081")
}
~~~

