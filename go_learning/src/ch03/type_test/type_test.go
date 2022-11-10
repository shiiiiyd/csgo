package type_test

import (
    "testing"
)

type MyInt int64

// Go 中只支持显示类型转换，别名也一样
func TestImplicit(t *testing.T) {
    var a int32 = 1
    var b int64
    b = int64(a)
    var c MyInt
    c = MyInt(b)
    t.Log(a, b, c)
}

// 指针
// Go支持指针但不支持运算
func TestPoint(t *testing.T) {
    a := 1
    aPtr := &a
    // aPtr = *aPtr + 1 // Go 中指针不支持运算
    t.Log(a, aPtr, *aPtr)
    t.Logf("%T, %T", a, aPtr)

}

// string
// string类型默认值是一个空值（""）
func TestString(t *testing.T) {
    // string 类型初始化是一个""（空值）
    var s string
    t.Log("@" + s + "@")
    if s == "" {
        t.Log("string类型初始化是一个\"\"值")
    }
}
