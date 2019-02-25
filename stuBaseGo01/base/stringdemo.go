package base

import (
	"fmt"
	"strconv"
	"strings"
)

func contiansTest() {
	s := "我在这"
	contains := strings.Contains(s, "1在")
	if contains {
		fmt.Println("找到")
	} else {
		fmt.Println("没有找到")

	}
}
func joinTst() {
	s := "123456"
	join := strings.Join(strings.Split(s, ""), "-")
	fmt.Print(join)
}
func otherTest() {
	//fold := strings.EqualFold("ab", "ab")
	fmt.Println("abcd"[:2])
	suffix := strings.TrimSuffix("abcde", "de")
	fmt.Println(suffix)
	replace := strings.Replace("abcdce", "c", "1", -1) //n 如果string 出现old有多次 你将决定替换几次 n=-1 屏蔽所有出现过的
	fmt.Println(replace)
	//repeat := strings.Repeat("1", -1) //重复拼接次数 count>=0
	//fmt.Println(repeat)
	//fmt.Print(strings.Compare("ab", "abc"))
}
func converStr() {
	//formatInt := strconv.FormatInt(123, 10) // 123 转成10进制的字符串
	formatInt2 := strconv.FormatInt(123, 2)
	fmt.Println(formatInt2)
	itoa := strconv.Itoa(1)
	fmt.Println(itoa)
	i, _ := strconv.Atoi("123") //字符串转整数
	fmt.Println(i)
	f, _ := strconv.ParseFloat("3.12569999999999999999999999699999999999999999996999999999999999999969999999999999999999699999999999999999996999999999999999999969999999999999999999999", 64)
	fmt.Println(f)
	strconv.AppendInt()
}
func main() {
	//contiansTest()
	//joinTst()
	converStr()
}
