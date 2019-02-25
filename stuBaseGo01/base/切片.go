package main

import (
	"fmt"
	"time"
	//"time"
)

func main01() {
	strings := make([]int, 0)
	var i int = 1
	for {
		strings = append(strings, i)
		i++
		if i == 10 {
			break
		}
	}

	fmt.Println(strings)      //[1 2 3 4 5 6 7 8 9]
	fmt.Println(strings[2:])  //[3 4 5 6 7 8 9]
	fmt.Println(strings[:5])  //[1 2 3 4 5]
	fmt.Println(strings[2:4]) //[3 4]
}

func main02() {
	var i int

	go func() {
		for {
			fmt.Println("go:", i)
		}
	}()
	for i < 100 {
		i++
		fmt.Println(i)
	}
	time.Sleep(20000)
}
func main() {
	//s := "aaaa"
	ints := make([]int, 3)
	ints[0] = 1
	ints[1] = 2
	ints[2] = 3
	fmt.Println(len(ints), cap(ints))
	ints = append(ints, 4)
	fmt.Println(len(ints), cap(ints))
}
