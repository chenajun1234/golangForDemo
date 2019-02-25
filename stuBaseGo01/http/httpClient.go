package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun/changdu")
	if err != nil {
		fmt.Println("http get err", err)
	}
	defer resp.Body.Close()
	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", bytes)

}
