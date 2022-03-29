package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",indexHandler) //路由绑定 indexHandler
	http.HandleFunc("/hello",helloHandler) //路由绑定 helloHandler
	log.Fatal(http.ListenAndServe(":9999",nil))//监听本地9999端口
}

//indexHandler 处理响应，访问/,响应 URL.Path=/
func indexHandler(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w,"URL.Path = %q\n", req.URL.Path)
}
//HelloHandler 处理响应，访问/hello,响应请求头的键值对信息
func helloHandler(w http.ResponseWriter, req *http.Request)  {
	for k,v := range req.Header{
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
}