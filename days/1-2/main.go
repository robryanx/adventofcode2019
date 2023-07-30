package main

import (
	"fmt"
	"math"

	"github.com/robryanx/adventofcode2019/modules/readinput"
)

func main() {
	total := 0
	for _, mass := range readinput.ReadInts("inputs/1/input.txt", "\n") {
		if mass > 0 {
			total += calculate_fuel(mass)
		}
	}

	fmt.Printf("%d\n", total)
}

func calculate_fuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3.0)) - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + calculate_fuel(fuel)
}
