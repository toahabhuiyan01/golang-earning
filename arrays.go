package main
import "fmt"
import "time"
func main(){
	// var a [10]int
	// a[3] = 42
	// i := a[3]

	// var a = [2]int{1, 2}
	// a := [2]int{1, 2}
	// a := [...]int{1, 2} //elipsis -> compiler figure out the size

	// //slices
	// var a []int
	// var a = []int [1, 2, 3, 4]

	// a := []int{1, 2, 3, 4}
	chars := []string{0:"a", 1:"b", 2:"c"}

	// var b = a[lo:hi]
	// var b = a[:3]
	// var b = a[1:3]

	// a  = make([]byte, 5, 5)

	x := [3]string{"toaha", "tanzel", "mahmud"}
	// s := x[:]

	for _, e := range x {
		fmt.Println(e)
	}
	for _, e := range chars {
		fmt.Println(e)
	}
	aaa := 0
	for range time.Tick(time.Second) {
		fmt.Println(aaa)
		aaa++
	}
}