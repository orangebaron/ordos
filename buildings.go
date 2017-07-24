package main

type basicBuilding struct {
  owner *player
  closed bool
}
func (b *basicBuilding) getOwner() *player { return b.owner }
func (b *basicBuilding) getClosed() bool { return b.closed }
func (b *basicBuilding) setClosed(c bool) { b.closed = c }
func (*basicBuilding) routinelyCalledFunc(int,int) {}
func (*basicBuilding) getInternalInt(string) int {return 0}
func (*basicBuilding) setInternalInt(string,int) {}

type apartment struct {
  basicBuilding
  improvementLvl int8
  occupancy int
  rate int
}
func (*apartment) getType() string { return "apartment" }
func (b *apartment) getRevenue() int { return b.occupancy * b.rate }
func (b *apartment) getRent() int { return [3]int{2,5,10}[b.improvementLvl] }
func (b *apartment) getInternalInt(name string) int {
  switch name {
  case "improvementLvl":
    return int(b.improvementLvl)
  case "occupancy":
    return b.occupancy
  case "rate":
    return b.rate
  default:
    return 0
  }
}
func (b *apartment) setInternalInt(name string,val int) {
  switch name {
  case "improvementLvl":
    b.improvementLvl = int8(val)
  case "occupancy":
    b.occupancy = val
  case "rate":
    b.rate = val
  }
}

type building interface {
  getType() string
  getOwner() *player
  getRevenue() int
  getRent() int
  getClosed() bool
  setClosed(bool)
  routinelyCalledFunc(x,y int)
  getInternalInt(string) int
  setInternalInt(string,int)
}
