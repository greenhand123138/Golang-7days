package main

import (
	"gee"
	"net/http"
)

func main()  {
	r:= gee.New()
	r.GET("/",func(c *gee.Context){
		c.HTML(http.StatusOK,"<h1>Hello Gee</h1>")
	})
	r.GET("/hello",func(c *gee.Context){
		c.String(http.StatusOK,"hello %s,you're at %s\n",c.Query("name"),c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK,gee.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")
}


/*
运行结果
chensongwen@chensongwendeMacBook-Pro ~ % curl -i http://localhost:9999
HTTP/1.1 200 OK
Context-Type: text/html
Date: Thu, 31 Mar 2022 08:41:25 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8

<h1>Hello Gee</h1>%

*/