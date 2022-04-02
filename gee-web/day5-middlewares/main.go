package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc  {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main()  {
	r:= gee.New()
	r.Use(gee.Logger())
	r.GET("/",func(c *gee.Context){
		c.HTML(http.StatusOK,"<h1>Hello Gee</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/",func(c *gee.Context){
			c.HTML(http.StatusOK,"<h1>Hello gee</h1>")
		})
		v1.GET("/hello", func(c *gee.Context){
			c.String(http.StatusOK,"hello %s,you're at %s\n",c.Query("name"),c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK,"hello %s, you're at %s\n",c.Param("name"),c.Path)
		})
	}
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