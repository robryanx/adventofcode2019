package intcode

import (
    "fmt"
    "strconv"
    "regexp"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Run_computer(computer_id int, codes []int, input <-chan int, result chan<- int, exit chan<- int) {
    regex := *regexp.MustCompile(`(?s)^(\d{0,1}?)(\d{0,1}?)(\d{0,1}?)(\d{1,2})$`)

    i := 0
    no_output := true
    last_output := -1

    relative_base := 0

    instruction_length := map[int]int {
        99: 0,
        1: 2,
        2: 2,
        3: 0,
        4: 1,
        5: 2,
        6: 2,
        7: 2,
        8: 2,
        9: 1,
    }

    opcode_loop:for {
        // parse the code for parameter mode
        instruction, parameter_modes := parse_opcode(codes[i], regex)

        parameters, base_parameters := parameter_check(codes, instruction_length[instruction], parameter_modes, i, relative_base)

        output_position := 0
        if i+instruction_length[instruction]+1 < len(codes) - 1 {
            output_position = codes[i+instruction_length[instruction]+1]
        }

        output_instruction(computer_id, i, output_position, instruction, parameter_modes, base_parameters, parameters)

        switch instruction {
            case 99:
                if(no_output) {
                    result <- codes[0]
                }
                exit <- last_output

                break opcode_loop;
            case 1:
                codes[codes[i+3]] = parameters[0] + parameters[1]
                i += 4
            case 2:
                codes[codes[i+3]] = parameters[0] * parameters[1]
                i += 4
            case 3:
                codes[codes[i+1]] = <- input
                i += 2
            case 4:
                fmt.Println(parameters[0])
                last_output = parameters[0]
                result <- parameters[0]
                no_output = false
                i += 2
            case 5:
                if(parameters[0] != 0) {
                    i = parameters[1]
                } else {
                    i += 3
                }
            case 6:
                if(parameters[0] == 0) {
                    i = parameters[1]
                } else {
                    i += 3
                }
            case 7:
                if(parameters[0] < parameters[1]) {
                    codes[codes[i+3]] = 1
                } else {
                    codes[codes[i+3]] = 0
                }

                i += 4
            case 8:
                if(parameters[0] == parameters[1]) {
                    codes[codes[i+3]] = 1
                } else {
                    codes[codes[i+3]] = 0
                }

                i += 4
            case 9:
                relative_base += parameters[0]

                i += 2
            default:
                fmt.Printf("invalid opcode %d\n", codes[i])
                break opcode_loop;
        }
    }
}

func parse_opcode(opcode int, regex regexp.Regexp) (int, [3]int) {
    var parameter_modes [3]int
    var instruction int

    opcode_string := strconv.Itoa(opcode)
    opcode_parts := regex.FindStringSubmatch(opcode_string)
    instruction, _ = strconv.Atoi(string(opcode_parts[4]))

    for j := 1; j<4; j++ {
        if(opcode_parts[j] == "") {
            opcode_parts[j] = "0"
        }

        value, err := strconv.Atoi(string(opcode_parts[j]))
        check(err)

        parameter_modes[3-j] = value
    }

    return instruction, parameter_modes
}

func parameter_check(codes []int, parameter_count int, parameter_modes [3]int, pointer int, relative_base int) ([]int, []int) {
    parameters := make([]int, parameter_count)
    base_parameters := make([]int, parameter_count)

    for i := 0; i<parameter_count; i++ {
        if parameter_modes[i] == 0 {
            parameters[i] = codes[codes[pointer+i+1]]
        } else if parameter_modes[i] == 1 {
            parameters[i] = codes[pointer+i+1]
        } else {
            parameters[i] = codes[relative_base+codes[pointer+i+1]]
            fmt.Println(relative_base+codes[pointer+i+1], relative_base, codes[pointer+i+1], parameters[i])
        }

        base_parameters[i] = codes[pointer+i+1]
    }

    return parameters, base_parameters
}

func output_instruction(computer_id int, position int, output_position int, instruction int, parameter_modes [3]int, base_parameters []int, parameters []int) {
    instruction_parts := []string{}
    for i := range base_parameters {
        if parameter_modes[2-i] == 0 || parameter_modes[2-i] == 2 {
            instruction_parts = append(instruction_parts, fmt.Sprintf("&%v (%v)", base_parameters[i], parameters[i]))
        } else {
            instruction_parts = append(instruction_parts, fmt.Sprintf("%v", base_parameters[i]))
        }
    }

    fmt.Printf("%v - pos: %v, out pos: %v, inst %v: %v\n", computer_id, position, output_position, instruction, strings.Join(instruction_parts, ", "))
}
