package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally")
	}()

	fmt.Println("Start")
	// panic 退出后不会执行除 defer代码的其他代码
	panic(errors.New("something wrong"))

	// os.Exit()退出后不会执行其他代码
	os.Exit(-1)
	fmt.Println("End")
}

// 将错误信息输出到日志中，但是没有解决改错误，如果出现的错误是核心资源耗尽的问题就会
// 形成僵尸服务进程，导致 health check 失败，就算强制恢复掉，健康的代码无法被检查
// 最好的方法是 "let it crash"恢复不确定性错误的最好方法
func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()

	panic(errors.New("wrong"))
}
