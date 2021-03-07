package domain

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Simulation struct {
	duration              int
	numberOfIntersections int
	numberOfStreets       int
	numberOfCars          int
	bonus                 int
	intersections         []Intersection
	streets               map[string]Street
	streetIntersectionMap map[string]int
	cars                  []Car
}

func NewSimulation(input []string) *Simulation {
	firstLine := strings.Split(input[0], " ")

	duration, _ := strconv.Atoi(firstLine[0])
	numberOfIntersections, _ := strconv.Atoi(firstLine[1])
	numberOfStreets, _ := strconv.Atoi(firstLine[2])
	numberOfCars, _ := strconv.Atoi(firstLine[3])
	bonus, _ := strconv.Atoi(firstLine[4])

	sim := Simulation{
		duration:              duration,
		numberOfIntersections: numberOfIntersections,
		numberOfStreets:       numberOfStreets,
		numberOfCars:          numberOfCars,
		bonus:                 bonus,
	}

	sim.addStreets(input[1 : 1+numberOfStreets])
	sim.addCars(input[1+numberOfStreets : len(input)-1])
	sim.mapIntersections()

	return &sim
}

func (s *Simulation) addStreets(input []string) {
	s.streets = make(map[string]Street)

	for _, line := range input {
		street := NewStreet(line)
		s.streets[street.name] = street
	}
}

func (s *Simulation) addCars(input []string) {
	s.cars = make([]Car, 0, len(input))

	for index, line := range input {
		s.cars = append(s.cars, NewCar(line, index))
	}
}

func (s *Simulation) mapIntersections() {
	s.intersections = make([]Intersection, s.numberOfIntersections)
	s.streetIntersectionMap = make(map[string]int)

	for _, street := range s.streets {
		s.intersections[street.endIntersectionId].id = street.endIntersectionId
		s.intersections[street.endIntersectionId].addStreet(street)

		s.streetIntersectionMap[street.name] = street.endIntersectionId
	}

	for index := 0; index < len(s.intersections); index++ {
		s.intersections[index].mapStreets()
	}
}

func (s *Simulation) OptimizeSchedule() {

	for _, car := range s.cars {
		for _, streetName := range car.route[:len(car.route)-1] {
			intersectionIndex := s.streetIntersectionMap[streetName]

			s.intersections[intersectionIndex].incrementPassingCarCounterFor(streetName)
		}
	}

	var totalUnusedStreets int
	var numberOfUnusedIntersections int

	for index := len(s.intersections) - 1; index >= 0; index-- {
		totalUnusedStreets += s.intersections[index].filterOutUnusedStreets()
		s.intersections[index].mapStreets()

		if len(s.intersections[index].streets) == 0 {
			s.intersections[index] = s.intersections[len(s.intersections)-1]
			s.intersections = s.intersections[:len(s.intersections)-1]
			numberOfUnusedIntersections++
		}
	}

	s.streetIntersectionMap = make(map[string]int)
	for index, intersection := range s.intersections {
		for _, street := range intersection.streets {
			s.streetIntersectionMap[street.name] = index
		}
	}

	fmt.Printf("Total number [%d] of unused streets.\nTotal number [%d] of unused intersections\n",
		totalUnusedStreets,
		numberOfUnusedIntersections)

	for i := 0; i < len(s.intersections); i++ {
		s.intersections[i].setCurrentSwitchLightTick()
	}
}

func (s *Simulation) Simulate() {
	fmt.Println("Start Simulation")

	for i := 0; i < s.duration; i++ {
		s.simulateCars(i)
		s.simulateIntersections(i)
	}

	fmt.Println("End Simulation")
}

func (s *Simulation) simulateIntersections(tick int) {
	for i := 0; i < len(s.intersections); i++ {
		s.intersections[i].simulateTick(tick)
	}
}

func (s *Simulation) simulateCars(tick int) {
	for index, car := range s.cars {
		intersectionIndex := s.streetIntersectionMap[car.getCurrentStreetName()]

		switch car.state {
		case Waiting:
			if s.intersections[intersectionIndex].canPassThroughIntersection(car) {
				s.intersections[intersectionIndex].moveCar(car)
				s.cars[index].move(tick, s.streets)
			}
		case Driving:
			if car.atEndOfStreet(tick) {
				if car.isAtFinish() {
					s.cars[index].finish(tick)
				} else if s.intersections[intersectionIndex].canPassThroughIntersection(car) {
					s.intersections[intersectionIndex].moveCar(car)
					s.cars[index].move(tick, s.streets)
				} else {
					s.cars[index].state = Waiting
					s.intersections[intersectionIndex].addCarToQueue(car)
				}
			}
		}
	}
}

func (s Simulation) CalculateScore() int {
	var score int

	for _, car := range s.cars {
		if car.state == Finished {
			score += s.bonus + (s.duration - car.finishedAt)
		}
	}

	return score
}

func (s Simulation) SaveSchedule(filename string) {
	file, err := os.Create("output/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(fmt.Sprintf("%d\n", len(s.intersections)))
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}

	for _, intersection := range s.intersections {
		_, err = writer.WriteString(intersection.toSubmissionFormat())
		if err != nil {
			log.Fatalf("Got erro while writing to a file. Err: %s", err.Error())
		}
	}

	_ = writer.Flush()
}

func (s Simulation) String() string {
	return fmt.Sprintf(
		"duration: [%d], numberOfIntersections: [%d], numberOfStreets: [%d], numberOfCars: [%d], bonus: [%d]\nintersections:\n%v\n streets:\n%v\n cars:\n%v",
		s.duration,
		s.numberOfIntersections,
		s.numberOfStreets,
		s.numberOfCars,
		s.bonus,
		s.intersections,
		s.streets,
		s.cars)
}
