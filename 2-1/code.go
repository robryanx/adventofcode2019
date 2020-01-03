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

    fmt.Printf("%f", codes)

    codes[1] = 12
    codes[2] = 2

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

    fmt.Printf("%d", codes[0]);
}
