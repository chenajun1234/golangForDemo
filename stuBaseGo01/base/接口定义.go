package base

import "fmt"

type demo10 interface {
	sayHello()
	show()
}
type person10 struct {
	id   int
	name string
	age  int
}
type student10 struct {
	id   int
	name string
	age  int
}
func (st *student10) sayHello(){
	fmt.Println("hello,我是学生")
}
func (p *person10) sayHello(){
	fmt.Println("hello,我是男人")
}
func (st *student10) show(){
	fmt.Println("我的名字是%s",st.name)
}
func (p *person10) show(){
	fmt.Println("我的名字是%s",p.name)
}
func main() {
	var de demo10;
	st:=student10{1,"aa",22}
	de=&st
	de.sayHello()
	p:=person10{2,"bb",33}
	de=&p
	de.sayHello()
}
