package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("mianji: ", g.area())
	fmt.Println("zhouchang: ", g.perim())
}

// ?
func detectCircle(g geometry) {
	if c1, ok := g.(circle); ok {
		fmt.Println("yuan:", c1.radius)
	}
}
func main() {
	r := rect{3, 4}
	c := circle{5}
	measure(r)
	fmt.Println("--")
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
