package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2019/modules/readinput"
)

func main() {
	numbers := readinput.ReadInts("inputs/8/input.txt", "")

	layers := chunk_ints(numbers, 150)

	var image [6][25]int

	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			pos := (y * 25) + x

			for layer := range layers {
				if layers[layer][pos] != 2 {
					image[y][x] = layers[layer][pos]
					break
				}
			}
		}
	}

	print_image(image)
}

func chunk_ints(numbers []int, chunk_size int) (chunks [][]int) {
	for chunk_size < len(numbers) {
		numbers, chunks = numbers[chunk_size:], append(chunks, numbers[0:chunk_size:chunk_size])
	}

	return append(chunks, numbers)
}

func print_image(image [6][25]int) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			if image[y][x] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%v", image[y][x])
			}
		}

		fmt.Printf("\n")
	}
}
