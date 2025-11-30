package main

import (
	"github.com/robryanx/adventofcode2019/util/intcode"
	"github.com/robryanx/adventofcode2019/util/readinput"
)

func main() {
	opcodes := readinput.ReadInts("inputs/11/input.txt", ",")

	input := make(chan int)
	result := make(chan int, 100)
	exit := make(chan int)

	go intcode.Run_computer(0, opcodes, input, result, exit, false)

	select {
	case <-exit:
		break
	}

	input <- 1

	<-exit
}
