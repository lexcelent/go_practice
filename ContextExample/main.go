package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrCanceled error = errors.New("canceled")

func rangeGenerator(cancel <-chan struct{}, start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func cancelChannelExample() {
	cancel := make(chan struct{})
	defer close(cancel)

	generated := rangeGenerator(cancel, 1, 10)
	for val := range generated {
		fmt.Println(val)
	}
}

// Параллельное объединение каналов
func mergeChannelsParallel(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	out := make(chan int)

	// Запуск горутин для чтения
	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for val := range channel {
				out <- val
			}
		}()
	}

	// Ожидание
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Ограничитель скорости.
// fn выполняется не более limit раз в секунду.
// handle - функция, которая выполняет fn с учетом лимита.
// cancel - останавливает ограничитель.
func rateLimiter(limit int, fn func()) (handle func() error, cancel func()) {
	done := make(chan struct{})
	requestsBuffer := make(chan struct{}, limit)

	limiter := time.Duration(1 * time.Second / time.Duration(limit))
	tick := time.Tick(limiter)

	handle = func() error {
		select {
		case <-done:
			return ErrCanceled
		default:
			<-tick
			go func() {
				requestsBuffer <- struct{}{}
				fn()
				<-requestsBuffer
			}()
			return nil
		}
	}

	cancel = func() {
		select {
		case <-done:
			return
		default:
			close(done)
		}
	}

	return handle, cancel
}

func main() {
	cancelChannelExample()
}
