package err

import (
	"errors"
	"testing"
)

// Go 语言中讲究快速失败原则，将发生错误的机制放置在其他代码前面
func GetFibonacci(n int) ([]int, error) {

	if n < 0 || n > 10 {
		return nil, errors.New("n should be in [0,10]")
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
