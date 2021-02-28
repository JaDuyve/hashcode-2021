package domain

import "fmt"

type Street struct {
	startIntersectionId int
	endIntersectionId   int
	name                string
	duration            int
}

func (s Street) String() string {
	return fmt.Sprintf("startIntersectionId: [%d], endIntersectionId: [%d], name: [%s], duration: [%d]\n",
		s.startIntersectionId,
		s.endIntersectionId,
		s.name,
		s.duration)
}
