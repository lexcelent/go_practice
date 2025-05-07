package shapes

// Треугольник
type Triangle struct {
	a, b, c int
}

func (t Triangle) Perimeter() int {
	return t.a + t.b + t.c
}

func NewTriangle(a, b, c int) *Triangle {
	return &Triangle{a, b, c}
}
