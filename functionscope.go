package main

import "fmt"

func main(){
	fmt.Println(outer())
	fmt.Println(scope())
}

func scope() func() int{
	outer_var := 2
	foo := func() int { return outer_var}
	return foo
}
// func another_scope() func() int {
// 	outer_var = 444 
// 	return foo
// }

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int{
		outer_var += 99
		return outer_var
	}
	return inner, outer_var
}