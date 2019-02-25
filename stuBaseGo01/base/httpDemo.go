package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/",&MyHandler{})
	http.ListenAndServe(":8888",mux)
}
//自定义handler
type MyHandler struct {

}
//实现handler接口
func ( *MyHandler) ServeHTTP(rsp http.ResponseWriter,r *http.Request)  {
	io.WriteString(rsp,"URL:"+r.URL.String())
}