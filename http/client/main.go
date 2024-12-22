package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: nil,
	}

	http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	req, _ := http.Get("http://localhost:8082")
	client.Get("http://localhost:8082")
	client.

	defer req.Body.Close()

	fmt.Println(req.Body)
}
