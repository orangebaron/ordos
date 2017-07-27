package main

import "./players"

var cpuPlayerList []players.ComputerPlayer
var networkPlayerList []players.NetworkPlayer
var playerList []players.Player

func main() {
	quitChan := make(chan struct{})
	setupServer(quitChan)
	gameLoop()
}
