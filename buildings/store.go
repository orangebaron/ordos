package buildings

type store struct {
	basicBuilding
	improvementLvl int8
	sales          int
	price          int
}

func (*store) GetType() string   { return "store" }
func (b *store) GetRevenue() int { return b.sales * b.price }
func (b *store) GetRent() int    { return [4]int{1, 2, 3, 5}[b.improvementLvl] }
func (b *store) GetInternalInt(name string) int {
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
func (b *store) SetInternalInt(name string, val int) {
	switch name {
	case "improvementLvl":
		b.improvementLvl = int8(val)
	case "sales":
		b.sales = val
	case "price":
		b.price = val
	}
}
func (b *store) ToString(x, y int) string {
	return "store," + string(b.improvementLvl) + "," + string(b.sales) + "," + string(b.price) + "," + b.basicBuilding.ToString(x, y)
}
