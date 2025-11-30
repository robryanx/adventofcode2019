package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2019/util/intcode"
	"github.com/robryanx/adventofcode2019/util/permutations"
	"github.com/robryanx/adventofcode2019/util/readinput"
)

func main() {
	opcodes := readinput.ReadInts("inputs/7/input.txt", ",")

	largest_signal := 0

	for phases := range permutations.Generate([]int{0, 1, 2, 3, 4}) {
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
		input := make(chan int)
		result := make(chan int)
		exit := make(chan int)

		run_opcodes := make([]int, len(opcodes))
		copy(run_opcodes, opcodes)

		go intcode.Run_computer(i, run_opcodes, input, result, exit, false)
		input <- phases[i]
		input <- output

		output = <-result

		<-exit
	}

	return output
}
