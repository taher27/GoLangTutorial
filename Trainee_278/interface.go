package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return float64(r.height * r.width)
}

func (c circle) area() float64 {
	return float64(c.radius * c.radius * math.Pi)
}
func (r rectangle) perim() float64 {
	return float64(2*r.height + 2*r.width)
}

func (c circle) perim() float64 {
	return float64(c.radius * math.Pi * 2)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Printf("Type: %T \n", g)
}

func main() {
	r := rectangle{3, 4}
	c := circle{5}
	fmt.Println("Rectangle: ")
	measure(r)
	fmt.Println("circle: ")
	measure(c)
}
