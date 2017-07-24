package main

import "time"
import "./buildings"

func gameLoop() {
	for {
		time.Sleep(1)
		bForEach(func(b *buildings.Building, x, y int) {
			(*b).RoutinelyCalledFunc(x, y)
		})
	}
}
