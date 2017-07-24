package players

type Player interface {
	getName() string
	getMoney() int
	addMoney(int)
}

type basicPlayer struct {
	name  string
	money int
}

func (p *basicPlayer) getName() string { return p.name }
func (p *basicPlayer) getMoney() int   { return p.money }
func (p *basicPlayer) addMoney(a int)  { p.money += a }
