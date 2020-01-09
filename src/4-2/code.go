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
        var adjacent_check byte
        decrease := false
        for j := 0; j < string_max - 1; j++ {
            if adjacent_check == byte(0) || adjacent_check != value_string[j] {
                if value_string[j] == value_string[j+1] && ((j == (string_max - 2)) || value_string[j+1] != value_string[j+2]) {
                    adjacent = true
                }

                adjacent_check = value_string[j]
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
            fmt.Println(i)

            passwords++
        }
    }

    fmt.Println(passwords)
}