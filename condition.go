package main
import "fmt"

func main(){
	c := 3
	d := 5
	
	if c > 5{
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}

	if a := c + d; a > 5{
		fmt.Println(a)
	}else{
		fmt.Println("not", a)
	}

	var val interface{}
	val = "foo"
	if str, ok := val.(string); ok {
		fmt.Println(str, ok)
	}else{
		fmt.Println("   ", str, ok)
	}
}