package main

import (
    "adventofcode/2019/modules/readinput"
    "adventofcode/2019/modules/intcode"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    opcodes := readinput.ReadInts("inputs/5/input.txt", ",")

    intcode.Run_computer(opcodes, []int{1})
}
