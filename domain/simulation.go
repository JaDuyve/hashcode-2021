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
	streets               []Street
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
	s.streets = make([]Street, 0, len(input))

	for _, line := range input {
		values := strings.Split(line, " ")
		startIntersectionId, _ := strconv.Atoi(values[0])
		endIntersectionId, _ := strconv.Atoi(values[1])
		duration, _ := strconv.Atoi(values[3])

		s.streets = append(s.streets, Street{
			startIntersectionId: startIntersectionId,
			endIntersectionId:   endIntersectionId,
			name:                values[2],
			duration:            duration,
		})
	}
}

func (s *Simulation) addCars(input []string) {
	s.cars = make([]Car, 0, len(input))

	for index, line := range input {
		values := strings.Split(line, " ")
		numberOfStreets, _ := strconv.Atoi(values[0])

		s.cars = append(s.cars, Car{
			id:                 index,
			numberOfStreets:    numberOfStreets,
			route:              values[1:],
			currentStreetIndex: 0,
			leaveStreetTick:    0,
		})
	}
}

func (s *Simulation) mapIntersections() {
	s.intersections = make([]Intersection, s.numberOfIntersections)

	for _, street := range s.streets {
		s.intersections[street.endIntersectionId].id = street.endIntersectionId
		s.intersections[street.endIntersectionId].addStreet(street.name)
	}
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
