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
	t := template.New("t.tmpl")

	_, err := t.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}

	name := "eintr"
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

func main() {
	http.HandleFunc("/", t)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err)
		return
	}
}
