package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2019/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    paths_str := readinput.ReadStrings("inputs/3/input.txt", "\n")
    paths := make([][][4]int, len(paths_str))

    for i := range paths_str {
        move_str := strings.Split(paths_str[i], ",")
        last := [4]int{0, 0, 0, 0}
        total_distance := 0

        for j := range move_str {
            runes := []rune(move_str[j])
            direction := string(runes[0])

            distance, err := strconv.Atoi(string(runes[1:]))
            check(err)

            next := last;
            next[2] = total_distance

            switch direction {
            case "L":
                for k:=(last[0]-1); k>=(last[0]-distance); k-- {
                    next[0] = k
                    next[3] = last[0] - k
                    paths[i] = append(paths[i], next)
                }
            case "R":
                for k:=(last[0]+1); k<=(last[0]+distance); k++ {
                    next[0] = k
                    next[3] = k - last[0]
                    paths[i] = append(paths[i], next)
                }
            case "U":
                for k:=(last[1]+1); k<=(last[1]+distance); k++ {
                    next[1] = k
                    next[3] = k - last[1]
                    paths[i] = append(paths[i], next)
                }
            case "D":
                for k:=(last[1]-1); k>=(last[1]-distance); k-- {
                    next[1] = k
                    next[3] = last[1] - k
                    paths[i] = append(paths[i], next)
                }
            }
            
            total_distance += distance
            last = next
        }
    }

    var shortest = -1
    for i := range paths[0] {
        for j := range paths[1] {
            if paths[0][i][0] == paths[1][j][0] && paths[0][i][1] == paths[1][j][1] {
                total_steps := paths[0][i][2] + paths[0][i][3] + paths[1][j][2] + paths[1][j][3]

                if shortest == -1 || total_steps < shortest {
                    shortest = total_steps

                    fmt.Println(i)
                    fmt.Println(paths[0][i])
                }
            }
        }
    }

    fmt.Println(shortest)
}