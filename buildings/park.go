package buildings

type park struct {
	basicBuilding
}

func (*park) GetType() string { return "park" }
func (*park) GetRent() int    { return 2 }
