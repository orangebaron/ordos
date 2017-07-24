package main

import "fmt"
import "./players"

var cpuPlayerList []players.ComputerPlayer
var networkPlayerList []players.NetworkPlayer
var playerList []players.Player

func main() {
	quitChan := make(chan struct{})
	networkPlayerList=[]players.NetworkPlayer{players.NewNetworkPlayer("uwe",3,"192.168.1.25")}
	serverFunc(quitChan)
	go gameLoop()
	fmt.Println("aaaaaaaaa")
	for {}
}
