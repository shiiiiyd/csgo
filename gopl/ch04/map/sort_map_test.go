package ch04_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortMap(t *testing.T) {

	ages := map[string]int{
		"alice":   21,
		"charlie": 18,
	}

	var names []string
	var value []int

	for name, v := range ages {
		names = append(names, name)
		value = append(value, v)
	}
	sort.Strings(names)
	// map 中的顺序是不确定的，使用 sort 排序。
	sort.Ints(value)
	fmt.Println(names)
	fmt.Println(value)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
