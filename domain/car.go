package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type State int

const (
	Driving  State = 0
	Waiting  State = 1
	Finished State = 2
)

type Car struct {
	id                 int
	numberOfStreets    int
	route              []string
	currentStreetIndex int
	leaveStreetTick    int
	finishedAt         int
	state              State
}

func NewCar(input string, id int) Car {
	values := strings.Split(input, " ")
	numberOfStreets, _ := strconv.Atoi(values[0])

	return Car{
		id:                 id,
		numberOfStreets:    numberOfStreets,
		route:              values[1:],
		currentStreetIndex: 0,
		leaveStreetTick:    0,
		finishedAt:         -1,
		state:              Waiting,
	}
}

func (c *Car) move(tick int, streets map[string]Street) {
	if c.numberOfStreets <= c.currentStreetIndex {
		c.finish(tick)
		return
	}

	c.currentStreetIndex++
	c.leaveStreetTick = tick + streets[c.getCurrentStreetName()].duration
	c.state = Driving
}

func (c Car) atEndOfStreet(tick int) bool {
	return c.leaveStreetTick == tick
}

func (c Car) getCurrentStreetName() string {
	return c.route[c.currentStreetIndex]
}

func (c *Car) finish(tick int) {
	c.state = Finished
	c.finishedAt = tick
}

func (c Car) String() string {
	return fmt.Sprintf(
		"id: [%d], numberOfStreets: [%d], currentStreetIndex: [%d], leaveStreetTick: [%d], finishedAt: [%d], route: [%s]\n",
		c.id,
		c.numberOfStreets,
		c.currentStreetIndex,
		c.leaveStreetTick,
		c.finishedAt,
		c.route)
}
