package main
import "fmt"

func main(){
	//no whlile or untill in golang
	for i := 1; i < 10; i++{
		fmt.Println(i)
	}

	fmt.Println()
	lst := []int{1, 4, 5, 7, 8, 0}
	for k, v := range lst{
		fmt.Println(k,"<-", v)
	}
}