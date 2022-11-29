package _select

import (
   "fmt"
   "testing"
   "time"
)

func service() string {
   time.Sleep(time.Millisecond * 100)
   return "Done"
}

func otherTask() {
   fmt.Println("working on something else")
   time.Sleep(time.Millisecond * 100)
   fmt.Println("task  is done.")
}

func TestService(t *testing.T) {
   fmt.Println(service())
   otherTask()
}

func AsyncService() chan string {
   //retCh := make(chan string)
   retCh := make(chan string, 1) // buffer chanel 更高效

   go func() {
      ret := service()
      fmt.Println("returned result.")
      retCh <- ret
      fmt.Println("service exited.")
   }()
   return retCh
}

func TestSelect(t *testing.T) {
   select {
   case ret := <-AsyncService():
      t.Log(ret)
   case <-time.After(time.Millisecond * 100):
      t.Error("time out")
   }
}
