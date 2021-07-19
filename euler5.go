package main

import (
    "fmt"
    "time"
)
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func main() {
    start := time.Now()
    res := smallestMultiple(20)
    fmt.Println(time.Since(start))
    fmt.Println(res)
}

func smallestMultiple(maxNum int) int {
    currentSmallest := 1
    for i := 1; i <= maxNum; i++{
        nextSmallest := currentSmallest
        for j := currentSmallest; !isMultiple(i, j); j+=currentSmallest{
            nextSmallest += currentSmallest
        }
        currentSmallest = nextSmallest
    }
    return currentSmallest
}

func isMultiple(highestNum int, num int) bool {
    for i := highestNum; i > 0; i--{
        if num % i != 0{
            return false
        }
    }
    return true
}
