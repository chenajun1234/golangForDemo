package base

import "fmt"

/**
@date 2018-10-11
 */
func main() {
	demo:=make([]interface{},3)
	demo[0]=123
	demo[1]=123.321
	demo[2]="字符串"
	for _,v:=range demo{
		if data,ok:=v.(int);ok{
			fmt.Println("整形",data)
		}else if data,ok:=v.(float64);ok{
			fmt.Println("浮点型",data)
		}else if data,ok:=v.(string);ok{
			fmt.Println("字符串",data)
		}
	}

}
