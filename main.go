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

	fmt.Print(sim)
	sim.SaveSchedule(filename)
}
