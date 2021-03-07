package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	. "trafic-signaling/domain"
)

func main() {
	files := [...]string{"a.txt", "b.txt", "c.txt", "d.txt", "e.txt", "f.txt"}

	for _, filename := range files {
		input, _ := ioutil.ReadFile("input/" + filename)
		sim := NewSimulation(strings.Split(string(input), "\n"))

		sim.OptimizeSchedule()
		sim.Simulate()
		fmt.Printf("Score: [%d]\n", sim.CalculateScore())

		sim.SaveSchedule(filename)
	}
}
