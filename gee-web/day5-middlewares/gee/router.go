package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parePattern(pattern string) []string  {
	vs := strings.Split(pattern,"/")
	parts := make([]string,0)
	for _, item := range vs{
		if item != ""{
			parts = append(parts,item)
			if item[0] == '*'{
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string,pattern string,handler HandlerFunc)  {
	log.Printf("Route %4s -%s", method,pattern)
	key := method + "-" + pattern
	parts := parePattern(pattern)
	_,ok:=r.roots[method]
	if !ok{
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern,parts,0)
	r.handlers[key] = handler
}

func (r *router)getRoute(method string,path string)(*node,map[string]string)  {
	searchParts := parePattern(path)
	params := make(map[string]string)
	root,ok:= r.roots[method]
	if !ok{
		return nil,nil
	}
	n := root.search(searchParts,0)
	if n !=nil {
		parts := parePattern(n.pattern)
		for index,part := range parts{
			if part[0] == ':'{
				params[path[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part)>1{
				params[part[1:]] = strings.Join(searchParts[index:],"/")
				break
			}
		}
		return n,params
	}
	return nil,nil
}

func (r *router) getRoutes(method string) []*node   {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node,0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context)  {
	n,params	:= r.getRoute(c.Method,c.Path)
	if n!= nil{
		key := c.Method + "-" +c.Path
		c.Partams=params
		c.handlers=append(c.handlers,r.handlers[key])
		}else{
			c.handlers=append(c.handlers, func(ctx *Context) {
				c.String(http.StatusNotFound,"404 not found:%s", c.Path)
			})
		}
	c.Next()
}

