package main

type basicPlayer struct {
	name  string
	money int
}

func (p *basicPlayer) getName() string { return p.name }
func (p *basicPlayer) getMoney() int   { return p.money }
func (p *basicPlayer) addMoney(a int)  { p.money += a }

type player interface {
	getName() string
	getMoney() int
	addMoney(int)
}
