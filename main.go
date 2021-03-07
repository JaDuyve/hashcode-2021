package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	. "trafic-signaling/domain"
)

func main() {
	files := [...]string{"a.txt", "b.txt", "c.txt", "d.txt", "e.txt", "f.txt"}
	var totalScore int

	for _, filename := range files {
		fmt.Printf("------------------------------------------------------\n--- Start file %s ---\n", filename)
		input, _ := ioutil.ReadFile("input/" + filename)
		sim := NewSimulation(strings.Split(string(input), "\n"))

		sim.OptimizeSchedule()
		sim.SaveSchedule(filename)

		sim.Simulate()
		score := sim.CalculateScore()
		totalScore += score
		fmt.Printf("File: %s, Score: [%d]\n", filename, score)

	}

	fmt.Printf("Total score: [%d]\n", totalScore)
}
