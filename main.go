package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	. "trafic-signaling/domain"
)

func main() {
	filename := "a.txt"
	input, _ := ioutil.ReadFile("input/" + filename)
	sim := NewSimulation(strings.Split(string(input), "\n"))

	sim.OptimizeSchedule()
	sim.Simulate()
	fmt.Printf("Score: [%d]\n", sim.CalculateScore())

	sim.SaveSchedule(filename)
}
