package main

import (
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
)

func main() {
    opcodes := readinput.ReadInts("inputs/9/input.txt", ",")
    memory := make([]int, 10000)

    opcodes = append(opcodes, memory...)

    input := make(chan int)
    result := make(chan int, 100)
    exit := make(chan int)

    go intcode.Run_computer(0, opcodes, input, result, exit)

    input <- 2

    <-exit
}