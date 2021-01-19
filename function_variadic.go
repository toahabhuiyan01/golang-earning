package main

import "fmt"

func main(){
	fmt.Println(adder(1,2,3))
	fmt.Println(adder(9, 9))
	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...))
}

func adder(args ...int) int {
	total := 0
	for k, v := range args{
		fmt.Println("    " , k)
		total += v
	}
	return total
}