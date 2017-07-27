package buildings

type park struct {
	basicBuilding
}

func (*park) GetType() string { return "park" }
func (*park) GetRent() int    { return 2 }
func (b *park) ToString(x, y int) string {
	return "park," + b.basicBuilding.ToString(x, y)
}
