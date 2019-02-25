package base

import "fmt"

func helio1p()  {
	fmt.Println("11111111110")
}
func helio1p2()  {
	//类似手动抛出一个异常
	panic("panic 异常")
}
func helio1p3()  {
	fmt.Println("333333333")
}
func main() {
	helio1p()
	helio1p2()
	helio1p3()
}
