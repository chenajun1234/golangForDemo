package base

import "fmt"

type people struct {
	id   int
	name string
	age  int
}
type students struct {
	man   people
	score float64
}
type student2 struct {
	people
	score float64
}

type student3 struct {
	//student3 不能嵌套本结构体  但指针可以
	*student2
	score float64
}

func main() {
	s:=student2{people:people{id:1,name:"aa",age:22},score:99}

	var s2 student2
	s2.score=99
	s2.id=11
	s2.name="bb"
	s2.age=33
	fmt.Println(s,s2)
}
