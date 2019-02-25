package base

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 1; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(<-ch)
	}
	time.Sleep(time.Minute)
}
