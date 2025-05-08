package ExampleSync

import (
	"fmt"
	"time"
)

// Канал завершения
func DoneChannelExample() {
	done := make(chan struct{})
	count := 4

	for range count {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Work has been done")
			// Передаем в канал сигнал о завершении
			done <- struct{}{}
		}()
	}

	// Ждем выполнения всех горутин
	for range count {
		<-done
	}
}
