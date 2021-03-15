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
