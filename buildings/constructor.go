package buildings

import "../players"

func NewBuilding(buildingType string, owner *players.Player) Building {
	switch buildingType {
	case "apartment":
		b := new(apartment)
		b.owner = owner
		return b
	case "store":
		b := new(store)
		b.owner = owner
		return b
	case "office":
		b := new(office)
		b.owner = owner
		return b
	case "park":
		b := new(park)
		b.owner = owner
		return b
	case "water":
		b := new(water)
		return b
	default:
		return nil
	}
}
