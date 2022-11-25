package loop

import (
	"testing"
)

func TestWhileLoop(t *testing.T) {
	for n := 1; n < 5; n++ {
		t.Log(n)
	}
}

func TestLoopWhile(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
