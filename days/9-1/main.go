package main

import (
	"github.com/robryanx/adventofcode2019/util/intcode"
	"github.com/robryanx/adventofcode2019/util/readinput"
)

func main() {
	opcodes := readinput.ReadInts("inputs/9/input.txt", ",")
	memory := make([]int, 10000)

	opcodes = append(opcodes, memory...)

	input := make(chan int)
	result := make(chan int, 100)
	exit := make(chan int)

	go intcode.Run_computer(0, opcodes, input, result, exit, false)

	input <- 1

	<-exit
}
