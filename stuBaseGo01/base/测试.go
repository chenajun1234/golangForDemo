package base

import "testing"

func add(a int, b int) int {
	return a + b
}
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{2, 5, 7},
		{6, 4, 10},
	}
	for _, v := range tests {
		if i := add(v.a, v.b); i != v.c {
			t.Error("add(%d,%d;)"+"got %d; expected %d",
				v.a, v.b, i, v.c)
		}
	}
}
