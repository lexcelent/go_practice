package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Простой GET-запрос
func simpleGetRequestExample() {
	fmt.Printf("Пример простого GET-запроса\n")
	const uri = "https://google.com/"
	resp, err := http.Get(uri)
	fmt.Println(resp.Status, err)
}

// Простой POST-запрос
func simplePostRequestExample() {
	fmt.Printf("Пример простого POST-запроса\n")
	const uri = "https://httpbingo.org/status/200"
	body := []byte("hello")
	resp, err := http.Post(uri, "text/plain", bytes.NewBuffer(body))
	fmt.Println(resp.Status, err)
}

// Запрос с параметрами в ссылке
func requestWithParams() {
	fmt.Printf("Пример запроса с параметрами \n")

	// Создаем клиента
	client := http.Client{Timeout: 3 * time.Second}

	// Создаем запрос
	const uri = "https://httpbingo.org/get"
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		panic(err)
	}

	// Наполняем запрос
	params := url.Values{}
	params.Add("id", "42")
	req.URL.RawQuery = params.Encode()

	// Редактируем заголовки
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Request-Id", "42")

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v %v \n", req.Method, req.URL)
	fmt.Println(resp.Status)
	fmt.Println()
}

func main() {
	simpleGetRequestExample()
	simplePostRequestExample()
	requestWithParams()
}
