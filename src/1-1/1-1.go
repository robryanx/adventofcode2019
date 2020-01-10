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
	file_raw, err := ioutil.ReadFile("../inputs/1/input.txt")
    check(err)
    lines := strings.Split(string(file_raw), "\n")

    total := 0;
    for i := 0; i<len(lines); i++ {
        mass, err := strconv.ParseFloat(lines[i], 64)
        if err == nil {
            total += int(math.Floor(mass / 3.0)) - 2
        }
    }

    fmt.Printf("%d\n", total);
}