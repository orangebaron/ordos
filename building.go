package main

type building interface {
	getType() string
	getOwner() *player
	getRevenue() int
	getRent() int
	getClosed() bool
	setClosed(bool)
	routinelyCalledFunc(x, y int)
	getInternalInt(string) int
	setInternalInt(string, int)
}

type basicBuilding struct {
	owner  *player
	closed bool
}

func (b *basicBuilding) getOwner() *player          { return b.owner }
func (*basicBuilding) getRevenue() int              { return 0 }
func (b *basicBuilding) getClosed() bool            { return b.closed }
func (b *basicBuilding) setClosed(c bool)           { b.closed = c }
func (*basicBuilding) routinelyCalledFunc(int, int) {}
func (*basicBuilding) getInternalInt(string) int    { return 0 }
func (*basicBuilding) setInternalInt(string, int)   {}
