package ch04_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 方式一
	ages01 := make(map[string]int)
	fmt.Println(ages01)

	// 方式二
	ages02 := map[string]int{
		"alice":   21,
		"charlie": 18,
	}
	// 相当于
	ages01["alice"] = 21
	ages01["charlie"] = 18
	fmt.Println(ages01)

	// 通过key访问
	ages02["alice"] = 19
	fmt.Println(ages02["alice"])
	fmt.Println(ages02)

	// 使用内置的 delete 函数删除元素
	delete(ages01, "alice")
	fmt.Println(ages01)

	// 查找不存在的值
	ages02["bob"] = ages02["bob"] + 1
	ages02["bob"] += 1
	fmt.Println(ages02)

	// 给定大小
	ages03 := make(map[string]int, 3)
	ages03["a"] = 95
	ages03["b"] = 96
	ages03["c"] = 97
	fmt.Println(ages03)

	ages03["d"] = ages03["c"] + 1
	ages03["e"] = ages03["d"] + 1
	fmt.Println(ages03)
}
