package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", say)
	http.ListenAndServe(":8888", nil)
}
func say(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("hello"))
	fmt.Println(req.Host)

}
