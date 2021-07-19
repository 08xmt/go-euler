package main

import (
    "fmt"
    "time"
)
//The difference between the sum of the squares of the first ten natural numbers and the square of the sum is 2640.
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func main() {
    start := time.Now()
    res := squareDiff(100)
    fmt.Println(time.Since(start))
    fmt.Println(res)
}

func squareDiff(num int) int {
    squareSum := 0
    summedSquares := 0
    for i := 1; i <= num; i++{
        squareSum += i*i
        summedSquares += i
    }
    return summedSquares*summedSquares-squareSum
}
