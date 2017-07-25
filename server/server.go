package main

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
	if r.URL.Path[len(r.URL.Path)-4:] == ".css" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	w.Write(b)
}
func eventFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("conn1")
	fmt.Println(r.RemoteAddr)
	plr := getPlayerOfIp(r.RemoteAddr)
	for len(plr.GetDataToBeSent()) == 0 {
		time.Sleep(time.Second / 2)
	}
	_, err := w.Write(plr.GetDataToBeSent()[0])
	if err == nil {
		plr.RemoveASentData()
	}
}
func chatFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("conn 2")
	fmt.Println(r.RemoteAddr)
	chat := ""
	r.ParseForm()
	for key := range r.Form {
		chat=key
		break
	}
	plr := getPlayerOfIp(r.RemoteAddr)
	data := []byte("CHAT"+plr.GetName()+": "+chat)
	for _, plr2 := range networkPlayerList {
		plr2.SendData(data)
	}
}
func setupServer(quitChan chan struct{}) {
	http.HandleFunc("/ordos.html",fileServe)
	http.HandleFunc("/ordos.js",fileServe)
	http.HandleFunc("/ordos.css",fileServe)
	http.HandleFunc("/event",eventFunc)
	http.HandleFunc("/chat",chatFunc)
	http.ListenAndServe(":8081",nil)
	fmt.Println("Server loaded")
}
