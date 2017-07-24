package main

import "fmt"

func main() {
	quitChan := make(chan struct{})
	go serverFunc(quitChan)
	fmt.Println("aaaaaaaaa")
	close(quitChan)
}
