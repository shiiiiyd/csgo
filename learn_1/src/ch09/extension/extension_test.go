package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	//dog := new(Dog).Pet
	//dog.Speak()
	//dog.Speak()
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

// 重写
//type Dog struct {
//    p *Pet
//}
//
//func (d *Dog) Speak() {
//	fmt.Print("Yao")
//	//d.p.Speak()
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}

func (d *Dog) Speak() {
	fmt.Print("Yao")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Dong")
}
