package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	h := Vertex{3, 4}
	h1 := Vertex{7, 5}
	h2 := Vertex{4, 6}
	fmt.Println(h.Abs())
	fmt.Println(h1.Abs())
	fmt.Println(h2.Abs())
}
