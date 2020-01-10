package main

import (
    "fmt"
    "math"
    "readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    total := 0
    for _, mass := range readinput.ReadFloats("../inputs/1/input.txt", "\n") {
        total += int(math.Floor(mass / 3.0)) - 2
    }

	fmt.Printf("%d\n", total);
}