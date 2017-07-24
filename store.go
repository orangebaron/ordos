package main

type store struct {
  basicBuilding
  improvementLvl int8
  sales int
  price int
}
func (*store) getType() string { return "store" }
func (b *store) getRevenue() int { return b.sales * b.price }
func (b *store) getRent() int { return [4]int{1,2,3,5}[b.improvementLvl] }
func (b *store) getInternalInt(name string) int {
  switch name {
  case "improvementLvl":
    return int(b.improvementLvl)
  case "sales":
    return b.sales
  case "price":
    return b.price
  default:
    return 0
  }
}
func (b *store) setInternalInt(name string,val int) {
  switch name {
  case "improvementLvl":
    b.improvementLvl = int8(val)
  case "sales":
    b.sales = val
  case "price":
    b.price = val
  }
}
