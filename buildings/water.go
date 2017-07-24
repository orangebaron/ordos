package buildings

import "../players"

type Water struct{}

func (*Water) GetType() string              { return "Water" }
func (*Water) GetOwner() *players.Player    { return nil }
func (*Water) GetRevenue() int              { return 0 }
func (*Water) GetRent() int                 { return 0 }
func (*Water) GetClosed() bool              { return false }
func (*Water) SetClosed(bool)               {}
func (*Water) RoutinelyCalledFunc(int, int) {}
func (*Water) GetInternalInt(string) int    { return 0 }
func (*Water) SetInternalInt(string, int)   {}
