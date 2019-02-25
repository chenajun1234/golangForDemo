package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬取网页内容
func DOGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 { //读取结束，或者，出问题
			//fmt.Println("resp.Body.Read err = ", e)
			break
		}

		result += string(buf[:n])
	}

	return
}
func DoSpider(i int, ch chan<- int) {
	//https://tieba.baidu.com/f?kw=%E5%89%91%E7%BD%913&ie=utf-8&pn=100
	//	url := "http://tieba.baidu.com/f?kw=%E5%89%91%E7%BD%913&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	url := "http://www.zhenai.com/zhenghun?1=1" + strconv.Itoa((i-1)*50)
	fmt.Println("url = ", url)
	//爬取内容
	result, err := DOGet(url) // DOGet(url)
	if err != nil {
		fmt.Println("DOGet err ", err)
		return
	}
	//写入文件
	fileName := strconv.Itoa(i) + ".html"
	file, e := os.Create(fileName)
	if e != nil {
		fmt.Println("os.Create err ", e)
		return
	}
	defer file.Close()
	file.WriteString(result)
	ch <- i
}
func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go DoSpider(i, ch)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-ch)
	}

}
func main() {
	var start, end int
	fmt.Print("请输入开始页:")
	fmt.Scan(&start)
	fmt.Print("请输入结束页:")
	fmt.Scan(&end)
	DoWork(start, end)

}
