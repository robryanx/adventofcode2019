package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2019/modules/intcode"
	"github.com/robryanx/adventofcode2019/modules/readinput"
)

func main() {
	opcodes := readinput.ReadInts("inputs/2/input.txt", ",")

	opcodes[1] = 12
	opcodes[2] = 2

	input := make(chan int)
	result := make(chan int)
	exit := make(chan int)

	go intcode.Run_computer(0, opcodes, input, result, exit, false)

	fmt.Printf("%d\n", <-result)
}
