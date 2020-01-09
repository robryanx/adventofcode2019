package main

import (
    "fmt"
    //"io/ioutil"
    //"strings"
    "strconv"
    //"math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input_min := 271973
    input_max := 785961

    passwords := 0

    for i := input_min; i < input_max; i++ {
        value_string := strconv.Itoa(i)
        string_max := len(value_string)

        adjacent := false
        decrease := false
        for j := 0; j < string_max - 1; j++ {
            if value_string[j] == value_string[j+1] {
                adjacent = true
            }

            value_a, err := strconv.Atoi(string(value_string[j]))
            check(err)

            value_b, err := strconv.Atoi(string(value_string[j+1]))
            check(err)

            if value_a > value_b {
                decrease = true
                break
            }
        }

        if adjacent && !decrease {
            passwords++
        }
    }

    fmt.Println(passwords)
}