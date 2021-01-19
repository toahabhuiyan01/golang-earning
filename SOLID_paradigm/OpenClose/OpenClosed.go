package main
import(
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type triangle struct {
	height float64
	base float64
}

func (t triangle) area() float64 {
	return t.height * t.base / 2
}

type shape interface {
	area() float64
}

type calcualtor struct {

}

func (a calcualtor) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 5}
	t := triangle{height: 5, base: 2}

	calc := calcualtor{}
	fmt.Println("area sum: ", calc.areaSum(c, s, t))
}