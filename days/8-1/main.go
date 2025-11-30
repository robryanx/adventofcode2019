package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2019/util/readinput"
)

func main() {
	numbers := readinput.ReadInts("inputs/8/input.txt", "")

	layers := chunk_ints(numbers, 150)

	min_count := [3]int{-1, 0, 0}

	for layer := range layers {
		count := [3]int{0, 0, 0}

		for number := range layers[layer] {
			count[layers[layer][number]]++
		}

		if min_count[0] == -1 || count[0] < min_count[0] {
			min_count = count
		}
	}

	fmt.Println(min_count[1] * min_count[2])
}

func chunk_ints(numbers []int, chunk_size int) (chunks [][]int) {
	for chunk_size < len(numbers) {
		numbers, chunks = numbers[chunk_size:], append(chunks, numbers[0:chunk_size:chunk_size])
	}

	return append(chunks, numbers)
}
