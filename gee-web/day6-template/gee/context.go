package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct{ // 定义属性
	Write http.ResponseWriter
	Req *http.Request
	Path string
	Partams map[string]string
	Method string
	StatusCode int
	handlers []HandlerFunc
	index 	int
	engine *Engine
}




func newContext(w http.ResponseWriter,req *http.Request) *Context {
	return &Context{
		Write: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
		index: -1,
	}
}

func (c *Context)Next()  {
	c.index++
	s:= len(c.handlers)
	for ; c.index <s; c.index++{
		c.handlers[c.index](c)
	}
}

func (c *Context)Fail(code int, err string)  {
	c.index = len(c.handlers)
	c.JSON(code,H{"message": err})
}

//提供了访问Param参数的方法
func (c *Context) Param(key string)string {
	value, _:= c.Partams[key]
	return value
}

//提供了访问PostForm参数的方法
func (c *Context) PostForm(key string)string  {
	return c.Req.FormValue(key)
}
//提供了访问Query参数的方法
func (c *Context)Query(key string)string  {
	return  c.Req.URL.Query().Get(key)
}

//转态码
func (c *Context)Status(code int)  {
	c.StatusCode=code
	c.Write.WriteHeader(code)
}

//请求头信息
func (c *Context)SetHeader(key string,value string)  {
	c.Write.Header().Set(key,value)
}

func (c *Context)String(code int,format string, values ...interface{})  {
		c.SetHeader("Content-Type", "text/plain")
		c.Status(code)
		c.Write.Write([]byte(fmt.Sprintf(format, values...)))
}

//提供对JSON类型的响应
func (c *Context)JSON(code int, obj interface{})  {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Write)
	if err:= encoder.Encode(obj); err!=nil{
		http.Error(c.Write,err.Error(),500)
	}
}

//提供对Data类型的响应
func (c *Context)Data(code int, data []byte)  {
	c.Status(code)
	c.Write.Write(data)
}
//提供对HTML类型的响应
func (c *Context)HTML(code int,name string,data interface{} )  {
	c.SetHeader("Context-Type","text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Write,name,data); err != nil{
		c.Fail(500,err.Error())
	}
}

