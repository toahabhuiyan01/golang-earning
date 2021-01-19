package main
import "fmt"
import "math"
//in golang there no class only struct
//here struct can have methods
//a struct can be a type . its also collection of fields

func main(){
	var v = Vertex{1, 2}
	fmt.Println(v.X)
	fmt.Println(v.Abs())

	//list of struct type
	var m = map[string]Vertex{
		"bell labs": {40, -74},
		"google": {37, -122},
	}
	fmt.Println(m)
	//pointers
	p := Vertex{1, 2}
	q := &p
	fmt.Println(q)
	var s *Vertex = new(Vertex)
}
type Vertex struct {
	X, Y int

}
// point := struct {
// 	X, Y int
// }{1, 2}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}




