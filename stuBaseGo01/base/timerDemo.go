package main

import (
	"fmt"
	"time"
)

func main01() {
	t := time.NewTimer(5 * time.Second)
	t.Reset(2 * time.Second) //重置
	//t.Stop()//关闭time而
	fmt.Println(time.Now())
	tc := t.C
	fmt.Println(<-tc)
}
func main() {
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	for {
		fmt.Println(i)
		<-ticker.C //没阻塞两秒
		i++
		if i == 5 {
			break
		}
	}
}
