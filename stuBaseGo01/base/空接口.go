package base

import "fmt"

func main() {
	var demo interface{}//相当万能指针
	demo=1
	fmt.Println(demo)
	demo="ad"
	fmt.Println(demo)
}
