package main

import (
	"fmt"
)

/*
Бинарный поиск в отсортированном массиве.
Вместо массива чисел можно попробовать сделать массив объектов.
Массив обязательно должен быть отсортирован
*/
func BinarySearch(sortedSlice []int, x int) (int, error) {
	if len(sortedSlice) == 0 {
		return -1, fmt.Errorf("элемент %d не найден", x)
	}

	if len(sortedSlice) == 1 {
		if sortedSlice[0] == x {
			return x, nil
		} else {
			return -1, fmt.Errorf("элемент %d не найден", x)
		}
	}

	middleIndex := len(sortedSlice)/2 + len(sortedSlice)%2

	if sortedSlice[middleIndex] == x {
		return x, nil
	}

	if sortedSlice[middleIndex] > x {
		return BinarySearch(sortedSlice[:middleIndex], x)
	} else {
		return BinarySearch(sortedSlice[middleIndex:], x)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(BinarySearch(numbers, 1))
}
