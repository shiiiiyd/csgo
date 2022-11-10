package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	//fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//temp := a
	//a = b
	//b = temp
	a, b = b, a
	t.Log(a, b)
}
