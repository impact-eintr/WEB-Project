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

自定义模板函数

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

避免XSS攻击

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


