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
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK,"heelo %s, you're at %s\n",c.Param("name"),c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})

		})
	}


	r.Run(":9999")
}


/*
运行结果
chensongwen@chensongwendeMacBook-Pro ~ % curl "http://localhost:9999/v1/hello?name=greenhand12138"
hello greenhand12138,you're at /v1/hello
chensongwen@chensongwendeMacBook-Pro ~ % curl "http://localhost:9999/v2/hello/greenhand12138"
404 not found:/v2/hello/greenhand12138
chensongwen@chensongwendeMacBook-Pro ~ %


*/