package main

import (
	"fmt"
	"sync"
	"time"
)

type atomic struct {
	a    int
	lock sync.Mutex
}

func (at *atomic) increment() {
	at.lock.Lock()
	defer at.lock.Unlock()
	at.a++

}
func (at *atomic) get() int {
	at.lock.Lock()
	defer at.lock.Unlock()
	return at.a
}
func main0001() {
	var a atomic
	a.increment()
	go func() {
		for i := 0; i < 100; i++ {
			a.increment()
			//fmt.Println(a.get())
		}
	}()

	time.Sleep(1000)
	fmt.Println(a.get())
}

func main() {
	nums := []int{2, 7, 12, 13,15,27}
	sum := twoSum(nums, 27)
	fmt.Printf("%d",sum)
}

func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums) &&i!=j ;j++  {
			if nums[i]+nums[j]==target{
				return []int{i, j}
			}
		}
	}
	return nil
}
