package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var codes []int;

    file_raw, err := ioutil.ReadFile("input.txt")
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

            if goal == run_computer(run_codes) {
                answer = (i*100) + j
            }
        }
    }

    fmt.Printf("%d", answer);
}

func run_computer(codes []int) int {
    opcode_loop:for i := 0; i<len(codes); i+=4 {
        fmt.Printf("%d - %d\n", i, codes[i]);

        switch codes[i] {
            case 99:
                break opcode_loop;
            case 1:
                codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
            case 2:
                codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
            default:
                fmt.Printf("invalid opcode %d\n", codes[i])
                break;
        }
    }

    return codes[0]
}