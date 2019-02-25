package base

import "fmt"

func main01() {
	fmt.Println("111111111111111")
	defer fmt.Println("2222222")//有点像finally
	defer fmt.Println("3333")
	fmt.Println("444")
}

func helloDefer1()  {
	fmt.Println("111111111111111")
}
func helloDefer2()  {
	defer fmt.Println("2222")
	fmt.Println("helloDefer2")
}
func main() {
	helloDefer1()
	//defer 函数执行完成时 在调用
	helloDefer2()
}