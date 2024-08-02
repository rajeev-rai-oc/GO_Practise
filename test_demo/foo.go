package main

import "strconv"

func Fooer(input int) string {

    isfoo := (input % 3) == 0

    if isfoo {
        return "Three"
    }

    return strconv.Itoa(input)
}