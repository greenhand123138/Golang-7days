package main

import (
	"gee"
	"net/http"
)



func main()  {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK,"Hello Greenhand12138\n")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"greenhand12138"}
		c.String(http.StatusOK,names[100])
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