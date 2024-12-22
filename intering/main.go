package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var memstat runtime.MemStats

func mem() {
	runtime.GC()
	runtime.ReadMemStats(&memstat)
	const MiB = 1024 * 1024
	fmt.Println("The program is now using", memstat.Alloc/MiB, "MiB")
}
func main() {

	const bookPath = "./dummy.txt"
	data, err := os.ReadFile(bookPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read", len(data), "bytes from", bookPath)
	mem()
	book := string(data)
	fmt.Println(len(book))
	//Bwords := findBwords(book)
	mem()
	// Use Bwords
	//fmt.Printf("The last B-word is %q\n", Bwords[len(Bwords)-1].Value())
	//mem()
}
