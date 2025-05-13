package main

import (
	"errors"
	"log"
	"time"
)

// Обработка с ожиданием.
// n-обработчиков работают с fn()
func withWorkers(n int, fn func()) (handle func(), wait func()) {
	// Обработчики будут брать токены из буферизованного канала
	tokens := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		tokens <- struct{}{}
	}

	// Каждый вызов handle() будет забирать один токен и
	// обрабатывать функцию
	handle = func() {
		<-tokens
		go func() {
			fn()
			tokens <- struct{}{}
		}()
	}

	// Функция ожидания
	wait = func() {
		for range n {
			<-tokens
		}
	}

	return handle, wait
}

// Обработка без ожидания.
// Т.е. здесь если все заняты, то функция об этом сообщит сразу
func withWorkersV2(n int, fn func()) (handle func() error, wait func()) {
	tokens := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		tokens <- struct{}{}
	}

	handle = func() error {
		select {
		case <-tokens:
			go func() {
				fn()
				tokens <- struct{}{}
			}()
			return nil
		default:
			return errors.New("busy")
		}
	}

	wait = func() {
		for range n {
			<-tokens
		}
	}

	return handle, wait
}

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	var result int

	done := make(chan struct{})
	go func() {
		result = fn()
		close(done)
	}()

	// <- done, если fn() успела до таймаута
	// <- time.After(), если прошло время
	select {
	case <-done:
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func Trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter to %s", msg)
	return func() { log.Printf("Exit from %s (%s)", msg, time.Since(start)) }
}

func main() {
	f := func() {
		time.Sleep(1 * time.Second)
	}

	handle, wait := withWorkers(2, f)
	defer Trace("main")()

	handle()
	handle()
	handle()
	handle()
	wait()

}
