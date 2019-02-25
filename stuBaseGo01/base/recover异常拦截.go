package base

import "fmt"

func recoverdemo(a int, b int) int {
	defer func() {
		fmt.Print(recover()) //打印出错误信息
	}()
	return a / b
}
func main01() {
	recoverdemo(10, 0) //runtime error: integer divide by zero
	fmt.Print("继续执行")
}
func recoverdemo02(a int, b int) int {
	defer func() {
		fmt.Println(recover()) //异常拦截一定要在异常放生前
	}()
	fmt.Println("111")
	var p *int
	*p = 123
	fmt.Println("111")
	return a / b
}
func main() {
	recoverdemo02(10, 0)
	fmt.Print("继续执行")
	/**
	111
	runtime error: invalid memory address or nil pointer dereference
	继续执行
	 */
}
