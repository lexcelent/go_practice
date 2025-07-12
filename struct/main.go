package main

import "fmt"

type Rectangle struct {
	a, b int
}

type Circle struct {
	r int
}

type ShapeCollection struct {
	rect   Rectangle
	circle Circle
}

func makeRectangle(a, b int) Rectangle {
	r := Rectangle{a, b}
	return r
}

func newRectangle(a, b int) *Rectangle {
	r := &Rectangle{a, b}
	return r
}

func (r Rectangle) area() int {
	return r.a * r.b
}

func main() {
	rect1 := Rectangle{2, 4}
	rect2 := Rectangle{a: 4, b: 7}
	rect3 := Rectangle{a: 4}

	fmt.Printf("%#v\t%#v\t%#v\n", rect1, rect2, rect3)

	collection := ShapeCollection{
		Rectangle{
			a: 4,
			b: 5,
		},
		Circle{
			r: 3,
		},
	}

	fmt.Printf("%#v\n", collection)
}
