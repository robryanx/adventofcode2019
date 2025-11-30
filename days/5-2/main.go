package main

import (
	"github.com/robryanx/adventofcode2019/util/intcode"
	"github.com/robryanx/adventofcode2019/util/readinput"
)

func main() {
	opcodes := readinput.ReadInts("inputs/5/input.txt", ",")

	input := make(chan int)
	result := make(chan int)
	exit := make(chan int)

	go intcode.Run_computer(0, opcodes, input, result, exit, false)
	input <- 5

	<-exit
}
