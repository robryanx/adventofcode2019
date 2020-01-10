package main

import (
    "fmt"
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

    file_raw, err := ioutil.ReadFile("../inputs/2/input.txt")
    check(err)
    opcodes_str := strings.Split(string(file_raw), ",")
    for i := 0; i<len(opcodes_str); i++ {
        opcode, err := strconv.Atoi(opcodes_str[i])
        check(err)

        codes = append(codes, opcode)
    }

    answer := 0
    goal := 19690720

    for i := 0; i<=99; i++ {
        for j := 0; j<=99; j++ {
            codes[1] = i
            codes[2] = j

            fmt.Printf("%f", codes);

            run_codes := make([]int, len(codes))
            copy(run_codes, codes)

            if goal == intcode.Run_computer(run_codes, 0) {
                answer = (i*100) + j
            }
        }
    }

    fmt.Printf("%d\n", answer);
}
