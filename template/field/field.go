package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string //大写导出的字段才会被渲染
	email    string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! you email is {{.email}}")
	p := Person{UserName: "Astaxie", email: "astaxie@gmail.com"}
	t.Execute(os.Stdout, p)
}
