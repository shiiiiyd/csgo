package _interface

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println(\"Hello World\")"
	//return `fmt.println("hello World")`
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
