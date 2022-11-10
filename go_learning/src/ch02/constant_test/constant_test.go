package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry(t *testing.T) {
	a := 1                                 //0001
	t.Log(Readable, Writeable, Executable) // 1,2,4
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
