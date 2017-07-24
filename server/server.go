package main

import "net" //now look at this net that I just found

func handleConn(conn net.Conn) {
	//do things
}

func serverFunc(quitChan chan struct{}) {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		select {
		case <-quitChan:
			return
		default:
			conn, _ := ln.Accept()
			go handleConn(conn)
		}
	}
}
