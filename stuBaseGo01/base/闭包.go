package base

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() IntGen {
	a, b := 0, 1
	return func() int64 {
		a, b = b, a+b
		return int64(a)
	}
}

type IntGen func() int64

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func pringtFileConet(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fibonacci()
	pringtFileConet(f)

}
