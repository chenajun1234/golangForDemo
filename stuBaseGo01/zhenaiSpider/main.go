package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	//resp, err := http.Get("https://www.liepin.com/zhaopin/?d_sfrom=search_fp_nvbar&init=1")
	resp, err := http.Get("https://tieba.baidu.com/f?kw=%E5%BE%B7%E4%BA%91%E8%89%B2&ie=utf-8&pn=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("请求失败")
		return
	}

	encode := determinEncoding(resp.Body)
	fmt.Println(encode)
	reader := transform.NewReader(resp.Body, encode.NewDecoder())
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	printArticleList(bytes)
	//fmt.Printf("%s\n", bytes)

}

func printArticleList(contents []byte) {
	compile := regexp.MustCompile(`<a rel="noreferrer" href="(/p/.+)" title=".+" target="_blank" class="j_th_tit ">(.+)</a>`)
	all := compile.FindAllSubmatch(contents, -1)
	for _, a := range all {
		fmt.Printf("title:%s URL:%s \n ", a[2], a[1])
	}
}
func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	encod, _, _ := charset.DetermineEncoding(bytes, "")
	return encod
}
