package string_func

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// 字符串分割
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}

	// 字符串连接
	t.Log(strings.Join(parts, "_"))
}

// 字符串转换
func TestStringConv(t *testing.T) {
	//数字转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)

	// 字符串转换称数字类型
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i + 10)
	}

}
