// 1.打开GC日志
// 只要在程序执行前加上环境变量 GODEBUG=gctrace=1，
// 例如：GODEBUG=gctrace=1 go test -bench=.
//
//	GODEBUG=gctrace=1 go run main.go
//
// 日志详细信息参考：https://godoc.org/runtime
// 2.使用 go tool trace
// go test -bench=. -trace trace.out
// 可视化 trace 信息：go tool trace.out
package gc_friendly

import "testing"

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

// 传值
func withValue(arr [NumOfElems]Content) int {
	// fmt.Println(&arr[2])
	return 0
}

// 传引用
func withReference(arr *[NumOfElems]Content) int {
	// b:=*arr
	// fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	// fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 值
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 引用传递
		withReference(&arr)
	}
	b.StopTimer()
}
