package main

import (
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
)

func main() {
    opcodes := readinput.ReadInts("inputs/5/input.txt", ",")

    input := make(chan int)
    result := make(chan int)
    exit := make(chan int)

    go intcode.Run_computer(0, opcodes, input, result, exit)
    input <- 1

    <- exit
}
