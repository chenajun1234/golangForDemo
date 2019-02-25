package main

import (
	"io"
	"net/http"
	"time"
)

//自定义handler
type MyHandler2 struct {

}
//实现handler接口
func ( *MyHandler2) ServeHTTP(rsp http.ResponseWriter,r *http.Request)  {
	io.WriteString(rsp,"URL:"+r.URL.String())
}
var muxMap map[string] func(http.ResponseWriter, *http.Request)
func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &MyHandler2{},
		ReadTimeout: 5 * time.Second,
	}
	muxMap= map[string]func(http.ResponseWriter, *http.Request){}
	muxMap["/hell"]=sayHello

	server.ListenAndServe()
}
func sayHello(rsp http.ResponseWriter,r *http.Request)  {
	rsp.Write([]byte(r.URL.String()+"hello"))
}