package empty_interface

import (
   "fmt"
   "testing"
)

// 类型断言
func DoSomething(p interface{}) {
   if i, ok := p.(int); ok {
      fmt.Println("Integer", i)
      return
   }
   if s, ok := p.(string); ok {
      fmt.Println("string", s)
      return
   }
   fmt.Println("Unknown Type")
}

func DoSomethingSwitch(p interface{}) {
   switch v := p.(type) {
   case int:
      fmt.Println("Integer", v)
   case string:
      fmt.Println("string", v)
   default:
      fmt.Println("Unknown Type")
   }
}

func TestEmptyInterfaceAssertion(t *testing.T) {
   DoSomething("10")
   DoSomething(10)
   DoSomethingSwitch("9")
   DoSomethingSwitch(9)
}

type A interface {
   ShowA() int
}

type B interface {
   ShowB() int
}

type Work struct {
   i int
}

func (w Work) ShowA() int {
   return w.i + 10
}

func (w Work) ShowB() int {
   return w.i + 20
}

func TestAssertion(t *testing.T) {
   var a A = Work{3}
   s := a.(Work)
   fmt.Println(s.ShowA())
   fmt.Println(s.ShowB())
}
