package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path, " hose: ", request.Host)

	fmt.Fprintf(w, "hello world")
}
func main() {

	http.HandleFunc("/", hello)
	//http.HandleFunc("/hello", nil)

	//err := http.ListenAndServe(":8082", nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	s := http.Server{
		Addr: ":8082",
	}
	fmt.Println(s)
}
