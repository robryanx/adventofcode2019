package main

import (
    "readinput"
    "intcode"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    opcodes := readinput.ReadInts("../inputs/5/input.txt", ",")

    intcode.Run_computer(opcodes, 1)
}
