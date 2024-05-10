package main

import (
	"fmt"
	"math"
)

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * (c.rad * c.rad)
}

type circle struct {
	rad float64
}

type rectangle struct {
	height float64
	width  float64
}

type from interface {
	area() float64
}

func writeArea(f from) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	r := rectangle{10, 15}
	writeArea(r) // 150
	c := circle{10}
	writeArea(c) // 314.16
}
