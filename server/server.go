package main

import "fmt"
import "strings"
import "net"  //now look at this net that I just found
//import "time" //time is a tool you can put on the wall or wear it on your wrist
//the past is far behind us, the future doesn't exist
import "./players"

func getPlayerOfIp(ip string) players.NetworkPlayer {
	for _, plr := range networkPlayerList {
		plrIp := plr.GetIp()
		if plrIp == ip[:strings.Index(ip, ":")] {
			return plr
		}
	}
	//for now, add a player to the list and return it
	networkPlayerList=append(networkPlayerList,players.NewNetworkPlayer("uwe",3,ip))
	return networkPlayerList[len(networkPlayerList)-1]
}

func handleConn1(conn net.Conn) {
	fmt.Println("conn 1")
	conn.Read([]byte{})
	_,err := conn.Write([]byte("aaa\x00"))
	fmt.Println(err)
	/*fmt.Println(conn.RemoteAddr().String())
	plr := getPlayerOfIp(conn.RemoteAddr().String())
	for len(plr.GetDataToBeSent()) == 0 {
		time.Sleep(time.Second / 2)
	}
	_, err := conn.Write(plr.GetDataToBeSent()[0])
	if err == nil {
		plr.RemoveASentData()
	}*/
}

func handleConn2(conn net.Conn) {
	fmt.Println("conn 2")
	fmt.Println(conn.RemoteAddr().String())
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

func checkLoop1(quitChan chan struct{}) {
	ln, _ := net.Listen("tcp", ":5252") //event listeners
	for {
		select {
		case <-quitChan:
			fmt.Println("65")
			return
		default:
			conn, _ := ln.Accept()
			go handleConn1(conn)
		}
	}
}
func checkLoop2(quitChan chan struct{}) {
	ln, _ := net.Listen("tcp", ":6565") //chat
	for {
		select {
		case <-quitChan:
			fmt.Println("65")
			return
		default:
			conn, _ := ln.Accept()
			go handleConn2(conn)
		}
	}
}

func setupServer(quitChan chan struct{}) {
	go checkLoop1(quitChan)
	go checkLoop2(quitChan)
	fmt.Println("server loaded")
}
