package main

import (
    "fmt"
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
)

func main() {
    opcodes := readinput.ReadInts("inputs/2/input.txt", ",")

    opcodes[1] = 12
    opcodes[2] = 2

    input := make(chan int)
    result := make(chan int)
    exit := make(chan int)

    go intcode.Run_computer(0, opcodes, input, result, exit)

    fmt.Printf("%d\n", <-result);
}
