package main

type park struct {
	basicBuilding
}

func (*park) getType() string { return "park" }
func (*park) getRent() int    { return 2 }
