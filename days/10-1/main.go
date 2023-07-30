package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2019/modules/readinput"
)

type position struct {
	x int
	y int
}

func main() {
	lines := readinput.ReadStrings("inputs/10/input.txt", "\n")

	position_list := []position{}

	for y, line := range lines {
		positions := strings.Split(line, "")
		for x, val := range positions {
			if val == "#" {
				position_list = append(position_list, position{
					x,
					y,
				})
			}
		}
	}

	most_seen := 0
	for _, pos := range position_list {
		seen := position_direction(position_list, pos.x, pos.y)
		if seen > most_seen {
			most_seen = seen
		}
	}

	fmt.Println(most_seen)
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func position_direction(position_list []position, x int, y int) int {
	seen := 0

	// calculate offsets
	offset_base := map[int]map[int]bool{}
	for _, pos := range position_list {
		if pos.x != x || pos.y != y {
			offset_x := pos.x - x
			offset_y := pos.y - y

			for i := max(abs(offset_x), abs(offset_y)); i > 0; i-- {
				if offset_x%i == 0 && offset_y%i == 0 {
					offset_x /= i
					offset_y /= i
					break
				}
			}

			if offset_base[offset_y] == nil {
				offset_base[offset_y] = make(map[int]bool)
			}

			if _, ok := offset_base[offset_y][offset_x]; !ok {
				offset_base[offset_y][offset_x] = true
				seen++
			}
		}
	}

	return seen
}
