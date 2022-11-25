package series

import "fmt"

// init 函数最先执行
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
