package buildings

import "../players"

type water struct{}

func (*water) GetType() string              { return "water" }
func (*water) GetOwner() *players.Player    { return nil }
func (*water) SetOwner(*players.Player)     {}
func (*water) GetRevenue() int              { return 0 }
func (*water) GetRent() int                 { return 0 }
func (*water) GetClosed() bool              { return false }
func (*water) SetClosed(bool)               {}
func (*water) RoutinelyCalledFunc(int, int) {}
func (*water) GetInternalInt(string) int    { return 0 }
func (*water) SetInternalInt(string, int)   {}
