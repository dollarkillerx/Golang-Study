package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak()  {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak()  {
	fmt.Println("this is DOg")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ",host)
}

func TestDog(t *testing.T)  {
	dog := new(Dog)
	dog.SpeakTo("Caho")
}