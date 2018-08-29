package main

import "fmt"

type shape interface {
	getArea() float64
}

type triagle struct {
	base   float64
	height float64
}
type square struct {
	sideLenght float64
}

func main() {
	t := triagle{base: 1.0, height: 2.0}
	s := square{sideLenght: 4.0}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triagle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}
