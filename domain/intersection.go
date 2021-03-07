package domain

import "fmt"

type Intersection struct {
	id                 int
	streets            []Street
	streetMap          map[string]int
	currentGreenStreet int
	switchLightTick    int
}

func (i *Intersection) addStreet(street Street) {
	street.greenLightDuration = 1
	i.streets = append(i.streets, street)
}

func (i *Intersection) simulateTick(tick int) {
	if i.switchLightTick == tick {
		i.currentGreenStreet = (i.currentGreenStreet + 1) % len(i.streets)
		i.switchLightTick = tick + i.streets[i.currentGreenStreet].greenLightDuration
	}
}

func (i *Intersection) mapStreets() {
	i.streetMap = make(map[string]int)

	for index, street := range i.streets {
		i.streetMap[street.name] = index
	}
}

func (i *Intersection) setCurrentSwitchLightTick() {
	i.switchLightTick = i.streets[i.currentGreenStreet].greenLightDuration
}

func (i *Intersection) moveCar(car Car) bool {
	return i.streets[i.currentGreenStreet].moveCarOut(car.id)
}

func (i Intersection) isLightGreenForCar(car Car) bool {
	return i.streetMap[car.getCurrentStreetName()] == i.currentGreenStreet
}

func (i *Intersection) addCarToQueue(car Car) {
	if streetIndex, exists := i.streetMap[car.getCurrentStreetName()]; exists {
		i.streets[streetIndex].addCar(car.id)
	}
}

func (i Intersection) String() string {
	return fmt.Sprintf(
		"id: [%d], numberOfIncomingStreets: [%d], outGoingStreets: [%v]\n",
		i.id,
		len(i.streets),
		i.streets)
}

func (i Intersection) toSubmissionFormat() string {
	output := fmt.Sprintf("%d\n%d\n",
		i.id,
		len(i.streets))

	for _, street := range i.streets {
		output += street.toSubmissionFormat()
	}

	return output
}
