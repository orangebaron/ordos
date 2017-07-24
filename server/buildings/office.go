package buildings

type office struct {
	basicBuilding
}

func (*office) GetType() string { return "office" }
func (*office) GetRent() int    { return 6 }
