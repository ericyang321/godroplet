package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Now make the shared methods `area()` and `perimeter()` for both structs
func (r *rectangle) area() float64 {
	return r.width * r.height
}
func (r *rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func findVals(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{width: 5.0, height: 10.0}
	c := circle{radius: 5.0}

	findVals(&r)
	findVals(&c)
}
