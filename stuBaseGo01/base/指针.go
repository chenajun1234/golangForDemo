package main

import "fmt"

func main01() {
	var a int = 10
	var s *int
	s = &a
	fmt.Println(s, *s)
	*s = 20
	fmt.Println(s, *s)
}
func main02() {
	var s *int
	// 内存地址0 到255系统占用 不允许读写
	*s = 122
	//s=x56464 //野指针 不能访问未知内存地址
	fmt.Println(s, *s)
}
func demos(p *int) {
	*p = 123456
}

//指针创建
func main() {
	var p *int = new(int)
	//p=
	fmt.Println(p, *p)
	*p = 213
	fmt.Println(p, *p)
	demos(p)
	fmt.Println(p, *p)
}
