package intcode

import (
    "fmt"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Run_computer(codes []int, input int) int {
    var parameter_modes [3]int
    var instruction int

    i := 0

    opcode_loop:for {
        fmt.Printf("%d - %d\n", i, codes[i]);

        // parse the code for parameter mode
        opcode_string := strconv.Itoa(codes[i])
        if(len(opcode_string) > 2) {
            if(len(opcode_string) == 3) {
                opcode_string = "00" + opcode_string
            } else if(len(opcode_string) == 4) {
                opcode_string = "0" + opcode_string
            }

            value, err := strconv.Atoi(opcode_string[len(opcode_string)-2:])
            check(err)
            instruction = value

            for j := 0; j<3; j++ {
                value, err := strconv.Atoi(string(opcode_string[j]))
                check(err)

                parameter_modes[2-j] = value
            }
        } else {
            instruction = codes[i]

            parameter_modes = [3]int{0, 0, 0}
        }

        switch instruction {
            case 99:
                break opcode_loop;
            case 1:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                codes[codes[i+3]] = parameters[0] + parameters[1]
                i += 4
            case 2:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                codes[codes[i+3]] = parameters[0] * parameters[1]
                i += 4
            case 3:
                codes[codes[i+1]] = input
                i += 2
            case 4:
                parameters := parameter_check(codes, 1, parameter_modes, i)

                fmt.Println(parameters[0])
                i += 2
            case 5:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                if(parameters[0] != 0) {
                    i = parameters[1]
                } else {
                    i += 3
                }
            case 6:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                if(parameters[0] == 0) {
                    i = parameters[1]
                } else {
                    i += 3
                }
            case 7:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                if(parameters[0] < parameters[1]) {
                    codes[codes[i+3]] = 1
                } else {
                    codes[codes[i+3]] = 0
                }

                i += 4
            case 8:
                parameters := parameter_check(codes, 2, parameter_modes, i)

                if(parameters[0] == parameters[1]) {
                    codes[codes[i+3]] = 1
                } else {
                    codes[codes[i+3]] = 0
                }

                i += 4
            default:
                fmt.Printf("invalid opcode %d\n", codes[i])
                break opcode_loop;
        }
    }

    return codes[0]
}

func parameter_check(codes []int, parameter_count int, parameter_modes [3]int, pointer int) ([]int) {
    parameters := make([]int, parameter_count)

    for i := 0; i<parameter_count; i++ {
        if parameter_modes[i] == 0 {
            parameters[i] = codes[codes[pointer+i+1]]
        } else {
            parameters[i] = codes[pointer+i+1]
        }
    }

    return parameters
}