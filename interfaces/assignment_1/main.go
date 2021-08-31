package main

import "fmt"

type shape interface{
	getArea() float64
}

type square struct{
	sideLength float64
}

type triangle struct{
	height float64
	base float64
}

func main() {
	square1 := square{sideLength: 4.4}

	triangle1 := triangle{height: 3.3, base: 2.33}
	
	printArea(square1)
	printArea(triangle1)

}

func (t triangle) getArea() float64  {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
