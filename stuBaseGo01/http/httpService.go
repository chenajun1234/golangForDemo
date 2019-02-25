package main

import (
	"fmt"
	"net/http"
)

func Myhander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.Host)
	fmt.Println(r.Body)
	fmt.Println(r.RequestURI)

	fmt.Fprint(w, "hello go")
}
func main() {
	http.HandleFunc("/go", Myhander)
	http.ListenAndServe(":8000", nil)
}
