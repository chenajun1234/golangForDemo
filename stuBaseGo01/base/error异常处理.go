package base

import (
	"errors"
	"fmt"
)

func demoError(a int ,b int ) (val int , errs error){
	if b==0{
		errs=errors.New("0不能作为除数")
		return
	}else{
		val=a/b
		return
	}
}

func main() {
	val, err := demoError(9, 0)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(val)
	}
}
