package main

type office struct {
	basicBuilding
}

func (*apartment) getType() string { return "office" }
func (b *apartment) getRent() int  { return 6 }
