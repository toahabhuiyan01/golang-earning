package main
import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) area() float64 {
	return math.Pow(c.length, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(objects ...object) float64{
	var sum float64
	for _, s := range objects {
		sum += s.area() + s.volume()
	}
	return sum
}
 
func main() {
	s1 := square{length: 3}
	s2 := square{length: 5}
	c1 := cube{square{length: 5}}
	c2 := cube{square{length: 4}}
	fmt.Println(areaSum(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum(c1, c2))
}

