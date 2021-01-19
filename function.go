package main

import "fmt"

func main() {
	//fmt.Println("Hello world!!")
	//var foo int
	//var foo int  = 42
	// var foo, bar int = 42, 1302
	//var foo := 42
	funcname()
	var x, s = returnmulti1()
	fmt.Println(x, s)

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	const constant = "This is a constant"
}

func funcname(){
	var x, str = returnmulti()
	fmt.Println(x, str)
	fmt.Println("this is new func")
	}
	
func returnmulti()(int, string){
	return 42, "hello"
	}

func returnmulti1()(n int, s string){
	n = 42
	s = "hello"

	return
}