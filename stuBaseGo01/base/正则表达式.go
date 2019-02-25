package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)
//const barListRe = `<a rel="noreferrer" href="(/p/.+)" [^>]*>([^<]+)</a>`
const barListRe = `<a data-field=.+ alog-group="p_author" class="p_author_name j_user_card" href="(/home/main.+)" [^>]*>(?s:(.+?))</a>`
func main() {
	//data-field="{&quot;un&quot;:&quot;chenajun1234&quot;,&quot;id&quot;:&quot;f7906368656e616a756e313233340a2c?t=1393314486&quot;}"
	//alog-group="p_author" class="p_author_name j_user_card"
	bytes, _ := ioutil.ReadFile("F:/work/gostu/src/stuBaseGo01/base/testRegexp.html")
	//fmt.Printf("%s \n",bytes)
	// + 表示任意字符  [a-zA-Z0-9]+@.+\..+  `<a rel="noreferrer" href="(/p/.+)" [^>]*>([^<]+)</a>`
	compile := regexp.MustCompile(barListRe)
	all := compile.FindAllSubmatch(bytes, -1)
	fmt.Println(len(all))
	for _,a:=range all{
		fmt.Printf("匹配了 %s \n 网址: %s \n 标题 %s",a[0],a[1],a[2])
	}
	//fmt.Println(len(all))

}
