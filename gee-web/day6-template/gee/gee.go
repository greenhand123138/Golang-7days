package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc 定义了 gee 使用的请求处理
type HandlerFunc func(ctx *Context)

// Engine实现ServeHTTP的接口
type (
	RouterGroup struct{
		prefix 			string
		middlewares		[]HandlerFunc //支持件
		parent 			*RouterGroup //支持嵌套
		engine			*Engine		//共享一个引擎实例
	}
	Engine struct{
		*RouterGroup
		router *router
		groups []*RouterGroup  //储存所有组
		htmlTemplates *template.Template
		funcMap 		template.FuncMap
	}
)


// new 是 gee.Engine 的构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup)Use(middlwares ...HandlerFunc )  {
	group.middlewares = append(group.middlewares,middlwares...)
}

func(group *RouterGroup) Group(prefix string) *RouterGroup{
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string,comp string, handler HandlerFunc)  {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s",method,pattern)
	group.engine.router.addRouter(method,pattern,handler)
}


//GET 定义了 添加GET 请求方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 定义了添加post请求方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) createStaticHandler(relativePath string,fs http.FileSystem)HandlerFunc  {
	absolutePath := path.Join(group.prefix, relativePath)
	fileSever := http.StripPrefix(absolutePath,http.FileServer(fs))
	return func(c *Context){
		file := c.Param("filepath")
		if _, err := fs.Open(file); err !=nil{
			c.Status(http.StatusNotFound)
			return
		}
		fileSever.ServeHTTP(c.Write,c.Req)
	}
}

// Staic方法暴雷给用户
func (group *RouterGroup) Static(relativePath string,root string)  {
	handler := group.createStaticHandler(relativePath,http.Dir(root))
	urlPattern := path.Join(relativePath,"/*fielpath")
	group.GET(urlPattern,handler)
}

func (engine *Engine)SetFuncMap(funcMap template.FuncMap)  {
	engine.funcMap = funcMap
}

func (engine *Engine)LoadHTMLGlob(pattern string)  {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

// Run 定义了启动 http 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}




func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups{
		if strings.HasPrefix(req.URL.Path,group.prefix){
			middlewares = append(middlewares,group.middlewares...)
		}
	}
	c := newContext(w,req)
	c.handlers=middlewares
	engine.router.handle(c)
}


