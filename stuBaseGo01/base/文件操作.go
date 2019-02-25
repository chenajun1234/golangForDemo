package base

import (
	"bufio"
	"fmt"

	"os"
)

func creteFileDemo() {
	f, e := os.Create("G:\\gostu\\fileSource\\a.text")
	if e != nil {
		fmt.Println("文件创建失败")
	}
	defer func() {
		f.Close()
	}()

	f.WriteString("文件创建成功")
}
func openFileDemo() {
	//f, e := os.OpenFile("G:\\gostu\\fileSource\\a.text", os.O_RDWR, 6)
	f, e := os.Open("G:\\gostu\\fileSource\\a.text")
	if e != nil {
		fmt.Println("文件打开失败")
	}
	defer func() {
		f.Close()
	}()
	b := make([]byte, 18)
	f.Read(b)
	fmt.Println(b)
	for _, value := range b {
		fmt.Printf("%c", value)
	}
}
func openFileDemo2() {
	f, e := os.OpenFile("G:\\gostu\\fileSource\\a.text", os.O_RDWR, 6)
	if e != nil {
		fmt.Println("文件打开失败")
	}
	defer f.Close()
	r := bufio.NewReader(f)
	//line, _, _ := r.ReadLine()
	//fmt.Println(string(line))
	//bytes, _ := r.ReadBytes('\n') //注意单引号
	b := make([]byte, 1)
	arrs := make([]byte, 18)
	//bytes, _ := r.Read(b)
	//fmt.Println(bytes)
	//fmt.Println(string(b))
	for {
		bytes, _ := r.Read(b)
		if bytes == 0 {
			break
		}
		fmt.Print(string(b))
	}
	fmt.Print(string(arrs))
}
func copyFile() {
	//f, e := os.Open("C:/Users/Administrator/Desktop/微信图片_20180925011032.png")
	f, e := os.Open("C:/Users/Administrator/Desktop/拷贝图片.png")
	cf, i := os.Create("C:/Users/Administrator/Desktop/拷贝图片2.png")

	if e != nil || i != nil {
		fmt.Println("文件打开失败")
	}
	defer func() {
		f.Close()
		cf.Close()
	}()
	bytes := make([]byte, 1024)
	for {
		n, _ := f.Read(bytes)
		cf.Write(bytes)
		if n == 0 {
			break
		}
	}

}
func main() {
	copyFile()
}
