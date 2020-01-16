package main

import (
    "fmt"
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
    "adventofcode/2019/modules/permutations"
)

func main() {
	opcodes := readinput.ReadInts("inputs/7/input.txt", ",")

	largest_signal := 0

	for phases := range permutations.Generate([]int{0, 1, 2, 3, 4}){
        signal := run_combination(opcodes, phases)

        if signal > largest_signal {
        	largest_signal = signal
        }
    }

    fmt.Println(largest_signal)
}

func run_combination(opcodes []int, phases []int) int {
	output := 0

	for i := range phases {
		output = intcode.Run_computer(opcodes, []int{phases[i], output})
	}

	return output
}