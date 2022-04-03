package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router  {
	r := newRouter()
	r.addRouter("GET","/",nil)
	r.addRouter("GET","/hello/:name",nil)
	r.addRouter("GET","/hello/b/c",nil)
	r.addRouter("GET","/hi/:name" ,nil)
	r.addRouter("GET","/assets/*filepath",nil)
	return r
}

func TestParsePattern(t *testing.T)  {
	ok	:= reflect.DeepEqual(parePattern("/p/:name"),[]string{"p","name"})
	ok = ok&& reflect.DeepEqual(parePattern("/p/*name"),[]string{"p","*"})
	ok = ok && reflect.DeepEqual(parePattern("/p/*name/*"),[]string{"p","*name"})
	if !ok{
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T)  {
	r := newTestRouter()
	n,ps := r.getRoute("GET","/hello/greenhand12138")

	if n == nil{
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name"{
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "greenhand12138" {
		t.Fatal("name should be equal to 'greenhand12138'")
	}
	fmt.Printf("matched path: %s, params['name']: %s", n.pattern, ps["name"])
}