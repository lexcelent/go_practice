package main

import (
	"fmt"
	waitgroup "goroutines/wait_group"
	"time"
)

func ChannelsExampleUsage() {
	c := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		// Отправляем значение в канал с задержкой
		c <- 5
	}()

	// Блокировка пока не будет получено значение из канала
	fmt.Printf("Число: %d", <-c)
}

func ChannelsExampleClose1() {
	in := make(chan int)

	go func() {
		for i := 1; i < 3; i++ {
			in <- i
		}
		close(in)
	}()

	res, ok := <-in
	fmt.Printf("Число в канале: %d, статус канала: %v\n", res, ok)

	res, ok = <-in
	fmt.Printf("Число в канале: %d, статус канала: %v\n", res, ok)

	res, ok = <-in
	fmt.Printf("Число в канале: %d, статус канала: %v\n", res, ok)

	res, ok = <-in
	fmt.Printf("Число в канале: %d, статус канала: %v\n", res, ok)
}

func ChannelsExampleClose2() {
	in := make(chan int)

	go func() {
		for i := 1; i < 3; i++ {
			in <- i
		}
		close(in)
	}()

	for res := range in {
		fmt.Printf("Число в канале: %d \n", res)
	}
}

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

func DeadlockExample() {
	out := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 1; i < 10; i++ {
			// Отправляется значение в канал и ожидается чтение
			out <- i
		}
		done <- struct{}{}
	}()

	// Ждем канал завершения. Блок
	<-done

	// Чтение из канала
	for number := range out {
		fmt.Println(number)
	}
}

// Пример с запуском нескольких короткоживущих горутин
func BufferedChannelExample1() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// N обработчиков = N горутин одновременно
	N := 2

	// Заполняем пул токенами
	pool := make(chan int, N)
	for i := 0; i < N; i++ {
		pool <- i
	}

	for _, num := range numbers {
		// Забираем токен и отдаем его горутине
		id := <-pool
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("Worker #%d: number %d \n", id, num)
			// Возвращаем токен
			pool <- id
		}()
	}

	// Ждем N обработчиков
	for range N {
		<-pool
	}
}

// Запускаются две долгие горутины, которые разгребают весь канал
func BufferedChannelExample2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	done := make(chan struct{})

	pending := make(chan int)
	N := 2

	// Отправляем все элементы в один канал
	go func() {
		for _, num := range numbers {
			pending <- num
		}
		close(pending)
	}()

	// Запускаем N обработчиков
	for i := 0; i < N; i++ {
		go func() {
			for num := range pending {
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker #%d: number %d \n", i, num)
			}
			done <- struct{}{}
		}()
	}

	// Ждем N обработчиков
	for range N {
		<-done
	}
}

func main() {
	waitgroup.UseWaitGroup()
	// waitgroup.CallOneByOne()
	// ChannelsExampleClose1()
	// ChannelsExampleClose2()
	// DoneChannelExample()
	// DeadlockExample()
	// BufferedChannelExample1()
	// BufferedChannelExample2()
}
