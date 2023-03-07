package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func main() {

}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (sh shape) printArea() {
	fmt.Println(sh.getArea())
}