package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/robryanx/adventofcode2019/modules/readinput"
)

type position struct {
	x         int
	y         int
	distance  int
	vaporised bool
}

type offsets struct {
	x         int
	y         int
	positions []position
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
					0,
					false,
				})
			}
		}
	}

	position := position_direction(position_list, 17, 23, 200)

	fmt.Println(position.x*100 + position.y)
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

func position_direction(position_list []position, x int, y int, count int) position {
	// calculate offsets
	offset_base := map[int]map[int][]position{}
	for _, pos := range position_list {
		if pos.x != x || pos.y != y {
			offset_x, base_offset_x := pos.x-x, pos.x-x
			offset_y, base_offset_y := pos.y-y, pos.y-y

			for i := max(abs(offset_x), abs(offset_y)); i > 0; i-- {
				if offset_x%i == 0 && offset_y%i == 0 {
					offset_x /= i
					offset_y /= i
					break
				}
			}

			if offset_base[offset_y] == nil {
				offset_base[offset_y] = make(map[int][]position)
			}

			pos.distance = abs(base_offset_x) + abs(base_offset_y)
			offset_base[offset_y][offset_x] = append(offset_base[offset_y][offset_x], pos)
		}
	}

	offsets_list := []offsets{}
	for y, offset_x := range offset_base {
		for x, positions := range offset_x {
			sort.Slice(positions, func(a, b int) bool {
				return positions[a].distance < positions[b].distance
			})

			offsets_list = append(offsets_list, offsets{
				x:         x,
				y:         y,
				positions: positions,
			})
		}
	}

	sort.Slice(offsets_list, func(a, b int) bool {
		quarter_a := quarter(offsets_list[a].x, offsets_list[a].y)
		quarter_b := quarter(offsets_list[b].x, offsets_list[b].y)

		if quarter_a > quarter_b {
			return true
		} else if quarter_a < quarter_b {
			return false
		}

		div_a := float64(offsets_list[a].x) / float64(offsets_list[a].y)
		div_b := float64(offsets_list[b].x) / float64(offsets_list[b].y)

		if div_a > div_b {
			return true
		}

		return false
	})

	total := 0
	for {
		found := false
		for i := range offsets_list {
			for j := range offsets_list[i].positions {
				pos := &offsets_list[i].positions[j]
				if !pos.vaporised {
					total++
					pos.vaporised = true

					if total == count {
						return *pos
					}

					found = true
					break
				}
			}
		}

		if !found {
			break
		}
	}

	return position{}
}

func quarter(x int, y int) int {
	switch true {
	case x >= 0 && y < 0:
		return 4
	case x >= 0 && y >= 0:
		return 3
	case x < 0 && y >= 0:
		return 2
	default:
		return 1
	}
}
