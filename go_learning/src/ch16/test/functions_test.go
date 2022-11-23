package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_square(t *testing.T) {
	type args struct {
		op int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 表格驱动测试
		{"01: ", args{1}, 1},
		{"02: ", args{2}, 4},
		{"03: ", args{3}, 9},
		{"04: ", args{4}, 16},
		{"05", args{5}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := square(tt.args.op); got != tt.want {
				t.Errorf("square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	want := [...]int{1, 4, 9}
	for i := 1; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != want[i] {
			t.Errorf("input is %d, want %d, actual %d", inputs[i], want[i], ret)
		}
	}
}

// Fail,Erro：该测试失败，但继续执行，其他测试继续执行
// FailNow，Fatal：该测试失败，该测试中止，其他测试继续执行
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Errorf("Error")
	fmt.Println("End")
}

func TestFailIncode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	want := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, want[i], ret)
	}
}
