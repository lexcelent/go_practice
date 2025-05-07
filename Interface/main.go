package main

import (
	"fmt"
	"interface/shapes"
)

// Интерфейс фигуры
type Shape interface {
	Perimeter() int
}

// Полиморфизм. На вход идет любая фигура, у которой реализован метод Perimeter() int
func PrintSomeInfo(shape Shape) int {
	fmt.Printf("Фигура: %T\n", shape)
	fmt.Printf("Параметры фигуры: %v\n", shape)
	fmt.Printf("Периметр фигуры: %d\n", shape.Perimeter())
	return 0
}

func main() {
	// Создаем фигуры
	rect := shapes.NewRectangle(4, 8)
	tr := shapes.NewTriangle(3, 4, 5)

	// Вызываем общий метод, в котором на вход передается фигура,
	// которая реализует интерфейс Shape
	PrintSomeInfo(rect)
	fmt.Println()

	PrintSomeInfo(tr)
	fmt.Println()

	// Вызываем метод Perimetr у Shape, передаем туда фигуру,
	// которая удовлетворяет интерфейсу
	Shape.Perimeter(rect)
	Shape.Perimeter(tr)
}
