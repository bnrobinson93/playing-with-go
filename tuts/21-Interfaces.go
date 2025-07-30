package main

import (
	"fmt"
	"math"
)

// interface collects methods
type geometry interface {
	area() float64
	perim() float64
}

// struct collects primitives
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 { // returns a function & that function returns a float64
	return r.width * r.height
}

func (r rect) perim() float64 { // as rect now returns both perim and area functions, it meets the creteria of the geometry interface
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // since it's an interface type, we know it'll have the functions defined
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
