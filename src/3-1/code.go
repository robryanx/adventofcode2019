package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file_raw, err := ioutil.ReadFile("../../inputs/3/input.txt")
    check(err)
    paths_str := strings.Split(string(file_raw), "\n")
    paths := make([][][2]int, len(paths_str))

    for i := range paths_str {
        move_str := strings.Split(paths_str[i], ",")
        last := [2]int{0, 0}

        for j := range move_str {
            runes := []rune(move_str[j])
            direction := string(runes[0])

            distance, err := strconv.Atoi(string(runes[1:]))
            check(err)

            next := last;

            switch direction {
            case "L":
                for k:=(last[0]-1); k>=(last[0]-distance); k-- {
                    next[0] = k
                    paths[i] = append(paths[i], next)
                }
            case "R":
                for k:=(last[0]+1); k<=(last[0]+distance); k++ {
                    next[0] = k
                    paths[i] = append(paths[i], next)
                }
            case "U":
                for k:=(last[1]+1); k<=(last[1]+distance); k++ {
                    next[1] = k
                    paths[i] = append(paths[i], next)
                }
            case "D":
                for k:=(last[1]-1); k>=(last[1]-distance); k-- {
                    next[1] = k
                    paths[i] = append(paths[i], next)
                }
            }

            last = next
        }
    }

    var shortest = -1
    for i := range paths[0] {
        for j := range paths[1] {
            if paths[0][i] == paths[1][j] {
                distance := int(math.Abs(float64(paths[0][i][0])) + math.Abs(float64(paths[0][i][1])))

                if shortest == -1 || distance < shortest {
                    shortest = distance

                    fmt.Println(i)
                    fmt.Println(paths[0][i])
                }
            }
        }
    }

    fmt.Println(shortest)
}