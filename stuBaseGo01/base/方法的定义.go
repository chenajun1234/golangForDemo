package base

import "fmt"

//给int类型取一个别名
type Int int
//给Int绑定方法
func (a Int) add(b Int) Int{
	return a+b
}
func main01() {
	var b Int=3
	b=b.add(5)
	fmt.Println(b)
}
type stu struct {
	id int
	name string
	age int
}
func (st stu) say(){
	fmt.Println(st)
}
func main02() {
	s:=stu{1,"张三",22}
	s.say()
}
func (st stu) editStu(name string ,age int){
	st.age=age
	st.name=name
	st.id=11
}
func (st *stu) editStu2(name string ,age int){
	st.age=age
	st.name=name
	st.id=11
}
func main() {
	var s stu
	s.editStu("李四",22)//值传递
	s.say()
	s.editStu2("李四2",22)//值传递
	s.say()
	//var s3 *stu//指针不能指向位置内存
	var s3 *stu=new(stu)
	s3.editStu2("李四3",22)//值传递
	s3.say()
}