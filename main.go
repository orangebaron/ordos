package main

import (
	"./players"
	"flag"
)

var cpuPlayerList []players.ComputerPlayer
var networkPlayerList []players.NetworkPlayer
var playerList []players.Player

func main() {
	flag.Parse()
	quitChan := make(chan struct{})
	setupServer(quitChan)
	gameLoop()
}
