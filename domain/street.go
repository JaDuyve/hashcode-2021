package domain

import (
	"fmt"
	"strconv"
	"strings"
	. "trafic-signaling/util"
)

type Street struct {
	startIntersectionId int
	endIntersectionId   int
	name                string
	duration            int
	greenLightDuration  int
	carQueue            Queue
	numberOfPassingCars int
}

func NewStreet(input string) Street {
	values := strings.Split(input, " ")
	startIntersectionId, _ := strconv.Atoi(values[0])
	endIntersectionId, _ := strconv.Atoi(values[1])
	duration, _ := strconv.Atoi(values[3])

	return Street{
		startIntersectionId: startIntersectionId,
		endIntersectionId:   endIntersectionId,
		name:                values[2],
		duration:            duration,
		carQueue:            NewQueue(),
		numberOfPassingCars: 0,
	}
}

func (s *Street) addCar(carId int) {
	s.carQueue.Add(carId)
}

func (s *Street) moveCarOut(carId int) {
	if s.carQueue.Length() != 0 && s.carQueue.Peek().(int) == carId {
		s.carQueue.Remove()
	}
}

func (s Street) isCarInFront(carId int) bool {

	return s.carQueue.Empty() || s.carQueue.Peek().(int) == carId
}

func (s Street) isQueueEmpty() bool {
	return s.carQueue.Empty()
}

func (s Street) String() string {
	return fmt.Sprintf("startIntersectionId: [%d], endIntersectionId: [%d], name: [%s], duration: [%d]\n",
		s.startIntersectionId,
		s.endIntersectionId,
		s.name,
		s.duration)
}

func (s Street) toSubmissionFormat() string {
	return fmt.Sprintf("%s %d\n",
		s.name,
		s.greenLightDuration)
}
