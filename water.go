package main

type water struct {
	basicBuilding
}

func (*water) getType() string { return "water" }
func (*water) getRent() int    { return 0 }
