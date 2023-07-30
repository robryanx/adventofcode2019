package intcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run_computer(computerId int, codes []int, input <-chan int, result chan<- int, exit chan<- int, debugOutput bool) {
	i := 0
	no_output := true
	last_output := -1

	relative_base := 0

	instruction_length := map[int]int{
		99: 0,
		1:  2,
		2:  2,
		3:  0,
		4:  1,
		5:  2,
		6:  2,
		7:  2,
		8:  2,
		9:  1,
	}

opcode_loop:
	for {
		// parse the code for parameter mode
		instruction, parameter_modes := parse_opcode(codes[i])

		parameters, base_parameters, output_position := parameter_check(codes, instruction_length[instruction], parameter_modes, i, relative_base)

		if debugOutput {
			output_instruction(computerId, i, output_position, instruction, parameter_modes, base_parameters, parameters)
		}

		switch instruction {
		case 99:
			if no_output {
				result <- codes[0]
			}
			exit <- last_output

			break opcode_loop
		case 1:
			codes[output_position] = parameters[0] + parameters[1]
			i += 4
		case 2:
			codes[output_position] = parameters[0] * parameters[1]
			i += 4
		case 3:
			codes[output_position] = <-input
			i += 2
		case 4:
			fmt.Println(parameters[0])
			last_output = parameters[0]
			result <- parameters[0]
			no_output = false
			i += 2
		case 5:
			if parameters[0] != 0 {
				i = parameters[1]
			} else {
				i += 3
			}
		case 6:
			if parameters[0] == 0 {
				i = parameters[1]
			} else {
				i += 3
			}
		case 7:
			if parameters[0] < parameters[1] {
				codes[output_position] = 1
			} else {
				codes[output_position] = 0
			}

			i += 4
		case 8:
			if parameters[0] == parameters[1] {
				codes[output_position] = 1
			} else {
				codes[output_position] = 0
			}

			i += 4
		case 9:
			relative_base += parameters[0]

			i += 2
		default:
			fmt.Printf("invalid opcode %d\n", codes[i])
			break opcode_loop
		}
	}
}

var opcodeRegex = regexp.MustCompile(`(?s)^(\d{0,1}?)(\d{0,1}?)(\d{0,1}?)(\d{1,2})$`)

func parse_opcode(opcode int) (int, [3]int) {
	var parameter_modes [3]int
	var instruction int

	opcode_string := strconv.Itoa(opcode)
	opcode_parts := opcodeRegex.FindStringSubmatch(opcode_string)
	instruction, _ = strconv.Atoi(string(opcode_parts[4]))

	for j := 1; j < 4; j++ {
		if opcode_parts[j] == "" {
			opcode_parts[j] = "0"
		}

		value, err := strconv.Atoi(string(opcode_parts[j]))
		check(err)

		parameter_modes[3-j] = value
	}

	return instruction, parameter_modes
}

func parameter_check(codes []int, parameter_count int, parameter_modes [3]int, pointer int, relative_base int) ([]int, []int, int) {
	parameters := make([]int, parameter_count)
	base_parameters := make([]int, parameter_count)
	output_position := 0

	for i := 0; i < parameter_count; i++ {
		if parameter_modes[i] == 0 {
			parameters[i] = codes[codes[pointer+i+1]]
		} else if parameter_modes[i] == 1 {
			parameters[i] = codes[pointer+i+1]
		} else {
			parameters[i] = codes[relative_base+codes[pointer+i+1]]
		}

		base_parameters[i] = codes[pointer+i+1]
	}

	if parameter_modes[parameter_count] == 0 {
		output_position = codes[pointer+parameter_count+1]
	} else if parameter_modes[parameter_count] == 1 {
		output_position = pointer + parameter_count + 1
	} else {
		output_position = relative_base + codes[pointer+parameter_count+1]
	}

	return parameters, base_parameters, output_position
}

func output_instruction(computerId int, position int, output_position int, instruction int, parameter_modes [3]int, base_parameters []int, parameters []int) {
	instruction_parts := []string{}
	for i := range base_parameters {
		if parameter_modes[2-i] == 0 || parameter_modes[2-i] == 2 {
			instruction_parts = append(instruction_parts, fmt.Sprintf("&%v (%v)", base_parameters[i], parameters[i]))
		} else {
			instruction_parts = append(instruction_parts, fmt.Sprintf("%v", base_parameters[i]))
		}
	}

	fmt.Printf("%v - pos: %v, out pos: %v, inst %v: %v\n", computerId, position, output_position, instruction, strings.Join(instruction_parts, ", "))
}
