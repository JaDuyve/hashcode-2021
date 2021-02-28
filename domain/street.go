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
	}
}

func (s *Street) addCar(carId int) {
	s.carQueue.Enqueue(carId)
}

func (s *Street) moveCarOut(carId int) bool {
	if s.isCarInFront(carId) {
		s.carQueue.Dequeue()
		return true
	}

	return false
}

func (s Street) isCarInFront(carId int) bool {
	return s.carQueue.Empty() || s.carQueue.Front() == carId
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
