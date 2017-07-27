package players

type Player interface {
	GetName() string
	GetMoney() int
	AddMoney(int)
}

type basicPlayer struct {
	name  string
	money int
}

func (p *basicPlayer) GetName() string { return p.name }
func (p *basicPlayer) GetMoney() int   { return p.money }
func (p *basicPlayer) AddMoney(a int)  { p.money += a }
