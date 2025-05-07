package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

// Функция, которая выполняется минимум 1 секунду
func DoSomeWork(str string) {
	time.Sleep(1 * time.Second)
	fmt.Println(str)
}

// Вызов с горутинами
func UseWaitGroup() {
	var wg sync.WaitGroup

	start := time.Now()
	for range 4 {
		wg.Add(1)
		go func() {
			DoSomeWork("Work has been done using goroutine")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Время с горутинами: %v \n", time.Since(start))
}

// Вызов по очереди
func CallOneByOne() {
	start := time.Now()
	for range 4 {
		DoSomeWork("Work has been done without goroutine")
	}
	fmt.Printf("Время без горутин: %v \n", time.Since(start))
}
