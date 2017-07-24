package buildings

type Park struct {
	basicBuilding
}

func (*Park) GetType() string { return "Park" }
func (*Park) GetRent() int    { return 2 }
