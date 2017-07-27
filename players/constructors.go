package players

func NewComputerPlayer(name string, money int) ComputerPlayer {
	return &computerPlayer{basicPlayer{name, money}} //add state variables later
}

func NewNetworkPlayer(name string, money int, ip string) NetworkPlayer {
	return &humanPlayer{basicPlayer{name, money}, ip, [][]byte{}}
}
