package main

import "fmt"

func main() {
	quitChan := make(chan struct{})
	go serverFunc(quitChan)
	go gameLoop()
	fmt.Println("aaaaaaaaa")
}
