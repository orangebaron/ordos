package main

import "fmt"
import "net"  //now look at this net that I just found
import "time" //time is a tool you can put on the wall or wear it on your wrist
//the past is far behind us, the future doesn't exist
import "./players"

func getPlayerOfIp(ip string) players.NetworkPlayer {
	for _, plr := range networkPlayerList {
		if plr.GetIp() == ip {
			return plr
		}
	}
	return nil
}

func handleConn1(conn net.Conn) {
	fmt.Println("conn 1")
	plr := getPlayerOfIp(conn.RemoteAddr().String())
	for len(plr.GetDataToBeSent()) == 0 {
		time.Sleep(time.Second / 2)
	}
	_, err := conn.Write(plr.GetDataToBeSent()[0])
	if err == nil {
		plr.RemoveASentData()
	}
}

func handleConn2(conn net.Conn) {
	fmt.Println("conn 2")
	var chat []byte
	_, err := conn.Read(chat)
	if err != nil {
		return
	}
	chat = append([]byte("CHAT"), chat...)
	plr := getPlayerOfIp(conn.RemoteAddr().String())
	for _, plr2 := range networkPlayerList {
		if plr2 != plr {
			plr2.SendData(chat)
		}
	}
}

func serverFunc(quitChan chan struct{}) {
	fmt.Println("not 65")
	ln1, _ := net.Listen("tcp", ":5252") //event listeners
	ln2, _ := net.Listen("tcp", ":6565") //chat
	for {
		select {
		case <-quitChan:
			fmt.Println("65")
			return
		default:
			conn1, _ := ln1.Accept()
			go handleConn1(conn1)
			conn2, _ := ln2.Accept()
			go handleConn2(conn2)
		}
	}
}
