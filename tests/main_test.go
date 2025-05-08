package main

import (
	"fmt"
	"testing"
)

// Test Setup + Test Teardown
// За это отвечает функция TestMain
func TestMain(m *testing.M) {
	// Исходные данные
	fmt.Println("Setup tests...")

	m.Run()

	fmt.Println("Teatdown tests...")
}

// Пример теста
func TestFirst(t *testing.T) {
	// Исходные данные
	got := GetSum(5, 6)
	want := 11

	// Проверка
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

// Пример теста с хардкодом
func TestFast(t *testing.T) {
	if GetSum(3, 4) != 7 {
		t.Errorf("Expected GetSum(3, 4) == 7")
	}
}

// Указать название теста
func TestFastWithName(t *testing.T) {
	t.Run("1Test TestGetSum(3,4)", func(t *testing.T) {
		if GetSum(3, 4) != 7 {
			t.Errorf("Ожидалось GetSum(3,4) == 7")
		}
	})

	t.Run("2Test TestGetSum(5,8)", func(t *testing.T) {
		if GetSum(5, 8) != 13 {
			t.Errorf("Ожидалось GetSum(5,8) == 13")
		}
	})
}

// Параметризация
func TestTableDriver(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{5, 6, 11},
		{3, 4, 7},
		{2, 11, 9},
		{0, 5, 5},
	}

	for _, test := range tests {

		// Название теста
		name := fmt.Sprintf("GetSum(%d,%d)", test.a, test.b)

		t.Run(name, func(t *testing.T) {
			got := GetSum(test.a, test.b)
			if got != test.want {
				t.Errorf("got %d; want %d", got, test.want)
			}
		})
	}
}
