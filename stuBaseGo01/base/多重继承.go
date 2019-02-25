package base

import "fmt"

type DemoA struct {
	id   int
	name string
	age  int
}
type DemoB struct {
	sex string
	age int
}
type DemoC struct {
	DemoA
	DemoB
	score float64
}

func main() {
	//d := DemoC{DemoA{1, "aa",12}, DemoB{"男", 22}, 55}
	var d DemoC;
	//d.age//父类字段同名 要加父类名称
	d.DemoA.age=3
	fmt.Println(d)
}
