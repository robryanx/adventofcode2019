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
        mass, err := strconv.Atoi(lines[i])

        if err == nil && mass > 0 {
            total += calculate_fuel(mass)
        }
    }

    fmt.Printf("%d", total);
}

func calculate_fuel(mass int) int {
    fuel := int(math.Floor(float64(mass) / 3.0)) - 2

    if fuel <= 0 {
        return 0;
    } else {
        return fuel + calculate_fuel(fuel)
    }
}