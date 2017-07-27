package main

import "./players"
import "./buildings"

func plrInit(p players.NetworkPlayer) {
	bForEach(func(b buildings.Building, x, y int) {
		p.SendData([]byte("BLDG" + b.ToString(x, y)))
	})
}
