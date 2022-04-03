package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)


type student struct {
	Name string
	Age  int8
}
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}




func main()  {
	r:= gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets","./static")
	stu1 := &student{Name: "greenhand12138", Age:22}
	stu2 := &student{Name: "CSW",Age: 20}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK,"css.tmpl",nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK,"arr.tmpl",gee.H{
			"title": "gee",
			"stuArr": [2]*student{stu1,stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK,"custom_func.tmpl",gee.H{
			"title": "gee",
			"now": time.Date(2022,4,3,14,10,0,0,time.UTC),
		})
	})



	r.Run(":9999")
}


/*
运行结果
chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhosts:9999/
curl: (6) Could not resolve host: localhosts
chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhost:9999/
<h1>Hello Gee</h1>%                                                             chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhost:9999/v2/hello/greenhand12138
{"message":"Internal Server Error"}



*/