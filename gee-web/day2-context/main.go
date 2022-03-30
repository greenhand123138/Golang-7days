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
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}


/*
运行结果
chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhost:9999
<h1>Hello Gee</h1>%                                                             chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhost:9999/hello
hello ,you're at /hello
chensongwen@chensongwendeMacBook-Pro ~ % curl http://localhost:9999/login
404 not found:/login
chensongwen@chensongwendeMacBook-Pro ~ % curl "http://localhost:9999/login" -X POST
{"password":"","username":""}
chensongwen@chensongwendeMacBook-Pro ~ % curl -i http://localhost:9999/
HTTP/1.1 200 OK
Context-Type: text/html
Date: Wed, 30 Mar 2022 13:17:44 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8

<h1>Hello Gee</h1>%                                                             chensongwen@chensongwendeMacBook-Pro ~ % curl "http://localhost:9999/login" -X POST -d "username=csw&password=20220330"
{"password":"20220330","username":"csw"}
chensongwen@chensongwendeMacBook-Pro ~
*/