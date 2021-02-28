package domain

import "fmt"

type Pair struct {
	a string
	b int
}

type Intersection struct {
	id                 int
	streets            []Pair
	currentGreenStreet int
	switchLightTick    int
}

func (i *Intersection) addStreet(streetName string) {
	i.streets = append(i.streets, Pair{streetName, 1})
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

	for _, pair := range i.streets {
		output += pair.toSubmissionFormat()
	}

	return output
}

func (p Pair) toSubmissionFormat() string {
	return fmt.Sprintf("%s %d\n",
		p.a,
		p.b)
}
