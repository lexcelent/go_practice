package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// Простой GET-запрос
func simpleGetRequestExample() {
	const uri = "https://google.com/"
	resp, err := http.Get(uri)
	fmt.Println(resp.Status, err)
}

// Простой POST-запрос
func simplePostRequestExample() {
	const uri = "https://httpbingo.org/status/200"
	body := []byte("hello")
	resp, err := http.Post(uri, "text/plain", bytes.NewBuffer(body))
	fmt.Println(resp.Status, err)
}

func main() {
	simpleGetRequestExample()
	simplePostRequestExample()
}
