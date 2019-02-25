package main

import (
	"fmt"
	"strconv"
)

type person struct {
	id int
	name string
	addr addr
}
type addr struct{
	addr string
}
type student struct {
	id int
	name string
	age  int
}
type student01 struct {
	person person
	age  int
}

func main() {
	s:=student01{
		person:&person{1,"aa",addr{"见面"}},
		age:23,
	}
	fmt.Print(s)
}

func main001() {
	s:=student{}
	s.id=1;
	s.age=23
	s.name="aa"
	fmt.Print(s)
	s2:=student{1,"bb",24}
	fmt.Print(s2)
	fmt.Print(s==s2)

}
func main002() {
	s:=[5]student{};
	for i:=0;i<len(s) ;i++  {
		temp:=student{}
		temp.id=i+1
		temp.name="name"+ strconv.Itoa(i+1)
		temp.age=10+i
		s[i]=temp
	}
	fmt.Println(s)
}
func main03() {
	var s [2]student;
	for i:=0;i<len(s) ;i++  {
		fmt.Scan(&s[i].id,&s[i].name,&s[i].age)
	}
	fmt.Println(s)
}
func demo(s map[string]student){
	s["三年一班"]=student{2,"bb",22}
	fmt.Println("s==",s)

}
func mai04n() {
	ss:= map[string]student{}
	ss["三年一班"]=student{1,"aa",11}
	demo(ss)
	fmt.Println(ss)

}