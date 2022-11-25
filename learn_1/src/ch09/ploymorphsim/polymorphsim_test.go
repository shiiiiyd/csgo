package ploymorphsim

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

func (j JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"hello world\")"
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T  %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	g := &GoProgrammer{}
	//g := new(GoProgrammer)
	java := new(JavaProgrammer)
	WriteFirstProgram(g)
	WriteFirstProgram(java)
}
