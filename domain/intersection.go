package domain

import "fmt"

type Intersection struct {
	id                 int
	streets            []Street
	streetMap          map[string]int
	currentGreenStreet int
	switchLightTick    int
	carPassed          bool
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

	i.carPassed = false
}

func (i *Intersection) mapStreets() {
	i.streetMap = make(map[string]int)

	for index, street := range i.streets {
		i.streetMap[street.name] = index
	}
}

func (i *Intersection) incrementPassingCarCounterFor(streetName string) {
	streetIndex := i.streetMap[streetName]
	i.streets[streetIndex].numberOfPassingCars++
}

func (i *Intersection) filterOutUnusedStreets() int {
	var numberOfUnusedStreet int
	for index := len(i.streets) - 1; index >= 0; index-- {
		if i.streets[index].numberOfPassingCars == 0 {
			i.streets[index] = i.streets[len(i.streets)-1]
			i.streets = i.streets[:len(i.streets)-1]
			numberOfUnusedStreet++
		}
	}

	return numberOfUnusedStreet
}

func (i *Intersection) setCurrentSwitchLightTick() {
	i.switchLightTick = i.streets[i.currentGreenStreet].greenLightDuration
}

func (i *Intersection) moveCar(car Car) {
	i.carPassed = true
	i.streets[i.currentGreenStreet].moveCarOut(car.id)
}

func (i Intersection) canPassThroughIntersection(car Car) bool {
	return i.streetMap[car.getCurrentStreetName()] == i.currentGreenStreet &&
		!i.carPassed &&
		i.streets[i.currentGreenStreet].isCarInFront(car.id)
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
