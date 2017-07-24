package buildings

import "../players"

func NewBuilding(buildingType string,owner *players.Player) Building {
  var b basicBuilding = nil
  switch buildingType {
  case "apartment":
    b=new(apartment)
  case "store":
    b=new(store)
  case "office":
    b=new(office)
  case "park"
    b=new(park)
  case "water":
    b=new(water)
  }
  b.owner=owner
  return b
}
