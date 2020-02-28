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

    series := []int{9, 8, 7, 6, 5}

    for phases := range permutations.Generate(series){
        signal := run_combination(opcodes, phases)

        if signal > largest_signal {
            largest_signal = signal
        }
    }

    fmt.Println(largest_signal)
}

func run_combination(opcodes []int, phases []int) int {
    output := 0

    channel := make([]chan int, len(phases))
    exit := make([]chan int, len(phases))

    for i := range phases {
        exit[i] = make(chan int)

        if i == 0 {
            channel[len(phases)-1] = make(chan int, 100)

            channel[len(phases)-1] <- phases[i]
            channel[len(phases)-1] <- 0
        } else {
            channel[i-1] = make(chan int, 100)

            channel[i-1] <- phases[i]
        }
    }

    for i := range phases {
        run_opcodes := make([]int, len(opcodes))
        copy(run_opcodes, opcodes)

        if i == 0 {
            go intcode.Run_computer(i, run_opcodes, channel[len(phases)-1], channel[i], exit[i])
        } else {
            go intcode.Run_computer(i, run_opcodes, channel[i-1], channel[i], exit[i])
        }

        fmt.Println(i)
    }
    
    output = <- exit[4]

    return output
}
