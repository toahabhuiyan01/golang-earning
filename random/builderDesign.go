package main
import (
	"fmt"
)

type mobile struct {
	name string
	price float64
}

func (m *mobile) setName(name string) *mobile {
	m.name = name
	return m
}

func (m *mobile) setPrice(price float64) *mobile {
	m.price = price
	return m
}

func (m mobile) getMobile() mobile {
	// return mobile{name: m.name, price: m.price}
	return m
}

func main(){
	m0 := mobile{}
	var m1 mobile
	fmt.Println(m0.setPrice(15000).setName("Samsung"))
	fmt.Println(m1.setPrice(5000).setName("Xiomi").getMobile())
}
