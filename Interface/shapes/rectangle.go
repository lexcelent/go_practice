package shapes

// Прямоугольник
type Rectangle struct {
	a, b int
}

func (r Rectangle) Perimeter() int {
	return r.a + r.a + r.b + r.b
}

func NewRectangle(a, b int) *Rectangle {
	return &Rectangle{a, b}
}
