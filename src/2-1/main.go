package main

import (
    "fmt"
    "readinput"
    "intcode"
)

func main() {
    opcodes := readinput.ReadInts("../inputs/2/input.txt", ",")

    opcodes[1] = 12
    opcodes[2] = 2

    result := intcode.Run_computer(opcodes, 0)

    fmt.Printf("%d\n", result);
}