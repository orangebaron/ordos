package buildings

type office struct {
	basicBuilding
}

func (*office) GetType() string { return "office" }
func (*office) GetRent() int    { return 6 }
func (b *office) ToString(x, y int) string {
	return "office," + b.basicBuilding.ToString(x, y)
}
