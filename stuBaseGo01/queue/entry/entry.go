package main

import (
	"fmt"
	"stuBaseGo01/queue"
)

func main() {
	q := queue.Queue{0}
	q.Push(1)
	q.Push(2)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
