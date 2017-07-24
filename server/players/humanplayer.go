package players

type NetworkPlayer interface {
	Player
	GetIp() string
	Kick()
	SendData([]byte)
}

type humanPlayer struct {
	basicPlayer
	ip string
}

func (p *humanPlayer) GetIp() string   { return p.ip }
func (p *humanPlayer) Kick()           {} //do this
func (p *humanPlayer) SendData([]byte) {} //do this
