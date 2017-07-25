package main

/*import "fmt"
import "strings"
import "net"  //now look at this net that I just found
import "time" //time is a tool you can put on the wall or wear it on your wrist
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
	time.Sleep(time.Second / 2)
	for {
		conn.Write([]byte{65})
	}
	fmt.Println(conn.RemoteAddr().String())
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
	fmt.Println(conn.RemoteAddr().String())
	var chat []byte
	_, err := conn.Read(chat)
	fmt.Println(chat,err)
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
*/




import "net/http"
import "fmt"
import "html"
import "time"
import "io/ioutil"
import "strings"
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
func fileServe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("conn0")
	b,e := ioutil.ReadFile("client/"+html.EscapeString(r.URL.Path)[1:])
	fmt.Println(e)
	w.Write(b)
}
func eventFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("conn1")
	//w.Write([]byte{65})
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	time.Sleep(0)
	/*fmt.Println(r.RemoteAddr)
	plr := getPlayerOfIp(r.RemoteAddr)
	for len(plr.GetDataToBeSent()) == 0 {
		time.Sleep(time.Second / 2)
	}
	_, err := w.Write(plr.GetDataToBeSent()[0])
	if err == nil {
		plr.RemoveASentData()
	}*/
}
func setupServer(quitChan chan struct{}) {
	http.HandleFunc("/ordos.html",fileServe)
	http.HandleFunc("/ordos.js",fileServe)
	http.HandleFunc("/ordos.css",fileServe)
	http.HandleFunc("/event",eventFunc)
	http.ListenAndServe(":8081",nil)
}
