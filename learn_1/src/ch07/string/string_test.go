package string_func_test

import "testing"

func TestStringToRune(t *testing.T) {
	s := "路漫漫其修远兮，吾将上下而求索"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestString(t *testing.T) {
	var s string
	// 初始化默认零值""
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3'         // string 是不可变的byte slice
	s = "\xE4\xB8\xAD" // 可以存储任何二进制数据
	// s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s)) // 是byte

	c := []rune(s)
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}
