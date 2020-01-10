package main

import (
    "io/ioutil"
    "strings"
    "strconv"
    "intcode"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var codes []int;

    file_raw, err := ioutil.ReadFile("../inputs/5/input.txt")
    check(err)
    opcodes_str := strings.Split(string(file_raw), ",")
    for i := 0; i<len(opcodes_str); i++ {
        opcode, err := strconv.Atoi(opcodes_str[i])
        check(err)

        codes = append(codes, opcode)
    }

    intcode.Run_computer(codes, 5)
}