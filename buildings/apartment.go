package buildings

type Apartment struct {
	basicBuilding
	improvementLvl int8
	occupancy      int
	rate           int
}

func (*Apartment) GetType() string   { return "Apartment" }
func (b *Apartment) GetRevenue() int { return b.occupancy * b.rate }
func (b *Apartment) GetRent() int    { return [3]int{2, 5, 10}[b.improvementLvl] }
func (b *Apartment) GetInternalInt(name string) int {
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
func (b *Apartment) SetInternalInt(name string, val int) {
	switch name {
	case "improvementLvl":
		b.improvementLvl = int8(val)
	case "occupancy":
		b.occupancy = val
	case "rate":
		b.rate = val
	}
}
