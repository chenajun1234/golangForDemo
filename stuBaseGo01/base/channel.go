package base

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, done chan bool) {
	for c := range ch { //循环接受channel的
		fmt.Printf("\n workid:%d val:%d", id, c)

		// 收完发完 打印完后就报错 原因
		//	// 1channel 都是阻塞的
		//	// 'a' 发完,worker收后 	done <- true 发 ,done等着接受
		//done <- true
		go func() { done <- true }()
	}
}
func creatChannel(id int) chans {
	ch := chans{
		ch:   make(chan int),
		done: make(chan bool),
	}
	go worker(id, ch.ch, ch.done)
	return ch
}

type chans struct {
	ch   chan int
	done chan bool
}

func channelDemo() {
	var ch [10]chans
	for i := 0; i < 10; i++ {
		ch[i] = creatChannel(i)
	}
	for i, w := range ch {
		w.ch <- 'a' + i
	}

	for i, w := range ch {
		w.ch <- 'A' + i
	}
	for _, w := range ch {
		<-w.done
		<-w.done
	}
	//

}
func bufferChannel() {
	ch := make(chan int, 3)
	ch <- 1 //给channel 输送值1
	ch <- 2 //给channel 输送值2
	ch <- 3 //给channel 输送值2
}
func main01() {
	channelDemo()
	//bufferChannel()
}

//---------------------------------------- sync.WaitGroup

func main02() {
	channelDemo2()
}

type chans2 struct {
	ch chan int
	wg *sync.WaitGroup
}

func creatChannel2(id int, wg *sync.WaitGroup) chans2 {
	ch := chans2{
		ch: make(chan int),
		wg: wg,
	}
	go worker2(id, ch.ch, ch.wg)
	return ch
}
func worker2(id int, ch chan int, wg *sync.WaitGroup) {
	for c := range ch { //循环接受channel的
		fmt.Printf("\n workid:%d val:%d", id, c)
		wg.Done()
	}
}
func channelDemo2() {
	var ch [10]chans2
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		ch[i] = creatChannel2(i, &wg)
	}
	wg.Add(20)
	for i, w := range ch {
		w.ch <- 'a' + i
	}

	for i, w := range ch {
		w.ch <- 'A' + i
	}
	wg.Wait()

}

//---------------------------------------- sync.WaitGroup 抽象

func main() {
	//channelDemo3()
	maps := map[string]string{
		"abc": "值",
	}
	fmt.Println(maps)
	m := make(map[string]int)
	m["abc"] = 123
	fmt.Println(m)
}

type chans3 struct {
	ch   chan int
	done func()
}

func creatChannel3(id int, wg *sync.WaitGroup) chans3 {
	ch := chans3{
		ch: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go worker3(id, ch)
	return ch
}
func worker3(id int, ch chans3) {
	for c := range ch.ch { //循环接受channel的
		fmt.Printf("\n workid:%d val:%d", id, c)
		ch.done()
	}
}
func channelDemo3() {
	var ch [10]chans3
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		ch[i] = creatChannel3(i, &wg)
	}
	wg.Add(20)
	for i, w := range ch {
		w.ch <- 'a' + i
	}

	for i, w := range ch {
		w.ch <- 'A' + i
	}
	wg.Wait()

}
