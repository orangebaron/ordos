package main

import "time"
import "./buildings"

func gameLoop() {
	x := 1 //so rent and revenue is paid every 10 seconds
	for {
		time.Sleep(1)

		bForEach(func(b buildings.Building, x, y int) {
			go b.RoutinelyCalledFunc(x, y)
		})

		for _, cpuPlayer := range cpuPlayerList {
			cpuPlayer.RoutinelyCalledFunc()
		}

		x++
		x %= 10
		if x == 0 {
			bForEach(func(b buildings.Building, x, y int) {
				go b.GetOwner().AddMoney(b.GetRevenue() - b.GetRent())
			})
		}
	}
}
