package players

type NetworkPlayer interface {
	Player
	GetIp() string
	Kick()
	SendData([]byte)
	GetDataToBeSent() [][]byte
	RemoveASentData()
}

type humanPlayer struct {
	basicPlayer
	ip           string
	dataToBeSent [][]byte
}

func (p *humanPlayer) GetIp() string             { return p.ip }
func (p *humanPlayer) Kick()                     {} //do this
func (p *humanPlayer) SendData(d []byte)         { p.dataToBeSent = append(p.dataToBeSent, d) }
func (p *humanPlayer) GetDataToBeSent() [][]byte { return p.dataToBeSent }
func (p *humanPlayer) RemoveASentData()          { p.dataToBeSent = p.dataToBeSent[1:] }
