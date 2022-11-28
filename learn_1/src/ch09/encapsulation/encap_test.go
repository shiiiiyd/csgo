package encapsulation

import (
   "fmt"
   "testing"
   "unsafe"
)

type Employee struct {
   Id   string
   Name string
   Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
   e0 := Employee{"0", "kob", 20}
   t.Log(e0)
   fmt.Printf("e0 Address is %x\n", unsafe.Pointer(&e0.Name))

   e1 := Employee{Name: "Mike", Age: 30}
   t.Log(e1)
   fmt.Printf("e1 Address is %x\n", unsafe.Pointer(&e1.Name))

   e2 := new(Employee) //返回指针
   e2.Id = "2"
   e2.Name = "Rose"
   e2.Age = 20
   t.Log(e2)
   fmt.Printf("e2 Address is %x\n", unsafe.Pointer(&e2.Name))

   e3 := &Employee{"3", "Jack", 20}
   t.Log(e3)
   fmt.Printf(" e3 Address is %x\n", unsafe.Pointer(&e3.Name))

   e4 := &Employee{"4", "dy", 18}
   t.Log(e4)
   fmt.Printf("e4 Address is %x\n", unsafe.Pointer(&e4.Name))

   e5 := &e2
   (*e5).Name = "test"
   fmt.Printf("e5 Address is %x\n", unsafe.Pointer(&(*e5).Name))
   t.Log(e2)

   e6 := &e0
   fmt.Printf("e6 address is %x\n", unsafe.Pointer(&e6.Name))

   t.Logf("e0 is %T", e0)
   t.Logf("e1 is %T", e1)
   t.Logf("e2 is %T", e2)
   t.Logf("e3 is %T", e3)
}

//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s >> Name:%s >> Age:%d", e.Id, e.Name, e.Age)
//}

func (e *Employee) String() string {
   fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
   return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
   e := Employee{"0", "Jack", 20}
   fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
   t.Log(e.String())
}
