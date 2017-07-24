package main

import "fmt"
import "./players"

var playerList []Player
func main() {
	quitChan := make(chan struct{})
	go serverFunc(quitChan)
	go gameLoop()
	fmt.Println("aaaaaaaaa")
}
