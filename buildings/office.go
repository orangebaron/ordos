package buildings

type Office struct {
	basicBuilding
}

func (*Office) GetType() string { return "office" }
func (*Office) GetRent() int    { return 6 }
