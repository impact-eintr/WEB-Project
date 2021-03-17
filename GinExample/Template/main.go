package main

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

func t(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t := template.New("t.tmpl")
	//解析模板
	_, err := t.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}

	name := "eintr"
	//执行模板
	t.Execute(w, name)
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

func main() {
	http.HandleFunc("/", t)
	http.HandleFunc("/xss", Xss)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err)
		return
	}
}
