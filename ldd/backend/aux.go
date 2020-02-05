package main

import (
	"strconv"
)

/* Auxiliar functions */
func strListToIntList(list []string) []int {

    var result []int

    for _, element := range list {
        i, _ := strconv.Atoi(element)
        result = append(result, i)
    }

    return result

}