package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	cl := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			cl <- i
			i++
		}
	}()
	return cl
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}
func main01() {
	var c1, c2 = generator(), generator()
	works := createWorker(0)
	b := false
	n := 0
	for {
		var active chan<- int
		if b {
			active = works
		}
		select {
		case n = <-c1:
			b = true
		case n = <-c2:
			b = true
		case active <- n:
			b = false

		}
	}

}
func fibonacc(cl chan<- int, flag <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case cl <- y:
			x, y = y, x+y
		case <-flag:
			return
		}
	}
}
func main() {
	cl := make(chan int)
	flag := make(chan bool)
	//消费者
	go func() {
		for i := 0; i < 8; i++ {
			fmt.Println(<-cl)
		}
		flag <- true
	}()
	//生产者
	fibonacc(cl, flag)
}
