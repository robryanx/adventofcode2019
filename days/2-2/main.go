package main

import (
    "fmt"
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
)

func main() {
    opcodes := readinput.ReadInts("inputs/2/input.txt", ",")

    answer := 0
    goal := 19690720

    for i := 0; i<=99; i++ {
        for j := 0; j<=99; j++ {
            opcodes[1] = i
            opcodes[2] = j

            run_opcodes := make([]int, len(opcodes))
            copy(run_opcodes, opcodes)

            input := make(chan int)
            result := make(chan int)
            exit := make(chan int)

            go intcode.Run_computer(0, run_opcodes, input, result, exit)

            if goal == <-result {
                answer = (i*100) + j
            }
        }
    }

    fmt.Printf("%d\n", answer);
}
