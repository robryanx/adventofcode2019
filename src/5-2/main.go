package main

import (
    "readinput"
    "intcode"
)

func main() {
    opcodes := readinput.ReadInts("../inputs/5/input.txt", ",")

    intcode.Run_computer(opcodes, 5)
}