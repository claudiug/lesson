package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	w float64
	h float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.h * s.w
}

func (s square) perim() float64 {
	return 2*s.w + 2*s.h
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	s := square{w: 10, h: 100}
	c := circle{radius: 10}
	measure(s)
	measure(c)
}
