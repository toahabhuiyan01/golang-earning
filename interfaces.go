package main
import "fmt"

func main(){
	type Awesomizer interface {
		Awesomize() string
	}
}

type Foo struct{}

func (foo, Foo) Awesomize() string {
	return "Awesome!"
}