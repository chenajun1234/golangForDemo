package main

import (
	"fmt"
	"gorutinePool/goPool"
	"sync"
	"time"
)

func worker(i, j, k int, wg *sync.WaitGroup) {
	m, n := 3, 5
	n, m = 5, 3
	fmt.Println(m, n)
	fmt.Printf("i:%d , j:%d, k:%d \n", i, j, k)
	wg.Done()
}

func main01() {
	var wg sync.WaitGroup
	start := time.Now()
	fmt.Println(start)

	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				wg.Add(1)
			}

		}
	}
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				go worker(i, j, k, &wg)
			}

		}
	}
	wg.Wait()
	sub := time.Now().Sub(start)
	fmt.Println("结束:", sub) //结束: 56.9414993s
}
func workPl() error {
	fmt.Println("无限的任务")
	return nil
}
func main() {
	start := time.Now()
	fmt.Println(start)
	task := goPool.Task{}
	task.NewTask(workPl)

	pool := goPool.Pool{}
	pool.NewPool(3)

	go func() {
		for {
			pool.EntryChannel <- task
		}

	}()
	pool.Run()
	sub := time.Now().Sub(start)
	fmt.Println("结束:", sub) //-结束: 13.2371727s

}
