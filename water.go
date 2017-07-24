package main

type water struct{}

func (*water) getType() string              { return "water" }
func (*water) getOwner() *player            { return nil }
func (*water) getRevenue() int              { return 0 }
func (*water) getRent() int                 { return 0 }
func (*water) getClosed() bool              { return false }
func (*water) setClosed(bool)               {}
func (*water) routinelyCalledFunc(int, int) {}
func (*water) getInternalInt(string) int    { return 0 }
func (*water) setInternalInt(string, int)   {}
