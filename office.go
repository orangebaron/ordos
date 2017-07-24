package main

type office struct {
	basicBuilding
}

func (*office) getType() string { return "office" }
func (*office) getRent() int    { return 6 }
