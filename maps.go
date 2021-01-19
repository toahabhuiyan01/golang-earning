package main
import "fmt"
import "Vertex"

func main(){
	// var m map[string]int
	m := make(map[string]int)
	m["toaha"] = 23
	fmt.Println(m["toaha"])

	// delete(m, "toaha")
	elem, ok := m["toaha"]
	fmt.Println(elem, ok)
}