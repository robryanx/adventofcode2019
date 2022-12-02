package main

import (
    "adventofcode/2019/modules/readinput"
    "fmt"
    "strings"
)

type position struct {
    x int
    y int
}

func main() {
    lines := readinput.ReadStrings("inputs/10/input.txt", "\n")

    field := make([][]bool, len(lines))
    position_list := []position{}

    for y, line := range lines {
        positions := strings.Split(line, "")
        field[y] = make([]bool, len(positions))
        for x, val := range positions {
            if val == "#" {
                field[y][x] = true
                position_list = append(position_list, position{
                    x,
                    y,
                })
            } else {
                field[y][x] = false
            }            
        }
    }

    most_seen := 0
    for _, pos := range position_list {
        seen := position_direction(field, position_list, pos.x, pos.y)
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

func position_direction(field [][]bool, position_list []position, x int, y int) int {
    y_max := len(field)
    x_max := len(field[0])

    seen := len(position_list) - 1

    // calculate offsets
    offset_list := []position{}
    for _, pos := range position_list {
        if pos.x != x || pos.y != y {
            offset_list = append(offset_list, position{
                x: pos.x - x,
                y: pos.y - y,
            })
        }
    }

    removed := map[int]map[int]bool{}

    for _, offset := range offset_list {
        x_check := x + offset.x
        y_check := y + offset.y

        for i:=max(abs(offset.x), abs(offset.y)); i>0; i-- {
            if offset.x % i == 0 && offset.y % i == 0 {
                offset.x /= i
                offset.y /= i
                break
            }
        }

        x_check += offset.x
        y_check += offset.y

        for x_check >= 0 && x_check < x_max && y_check >= 0 && y_check < y_max {
            if field[y_check][x_check] {
                if _, ok := removed[y_check][x_check]; !ok {
                    if removed[y_check] == nil {
                        removed[y_check] = make(map[int]bool)
                    }

                    removed[y_check][x_check] = true
                    seen--
                }
            }

            x_check += offset.x
            y_check += offset.y
        }
    }

    return seen
}
