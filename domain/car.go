package domain

import "fmt"

type Car struct {
	id                 int
	numberOfStreets    int
	route              []string
	currentStreetIndex int
	leaveStreetTick    int
}

func (c Car) String() string {
	return fmt.Sprintf(
		"id: [%d], numberOfStreets: [%d], currentStreetIndex: [%d], leaveStreetTick: [%d], route: [%s]\n",
		c.id,
		c.numberOfStreets,
		c.currentStreetIndex,
		c.leaveStreetTick,
		c.route)
}
