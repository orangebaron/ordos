package buildings

type Store struct {
	basicBuilding
	improvementLvl int8
	sales          int
	price          int
}

func (*Store) GetType() string   { return "Store" }
func (b *Store) GetRevenue() int { return b.sales * b.price }
func (b *Store) GetRent() int    { return [4]int{1, 2, 3, 5}[b.improvementLvl] }
func (b *Store) GetInternalInt(name string) int {
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
func (b *Store) SetInternalInt(name string, val int) {
	switch name {
	case "improvementLvl":
		b.improvementLvl = int8(val)
	case "sales":
		b.sales = val
	case "price":
		b.price = val
	}
}
