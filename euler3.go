package main

import "fmt"
import "time"
//The prime factors of 13195 are 5, 7, 13 and 29..
//What is the largest prime factor of the number 600851475143 ?.
func main() {
    start := time.Now()
    res := primeFact(600851475143)
    fmt.Println(time.Since(start))
    start = time.Now()
    res = primeFactSkip(600851475143)
    fmt.Println(time.Since(start))
    fmt.Println(res)
}

func primeFact(num int) int {
    smallestDivisor := 2
    for num % smallestDivisor != 0{
        smallestDivisor += 1
    }
    if(smallestDivisor == num){
        return smallestDivisor
    }
    return primeFact(num/smallestDivisor)
}
func primeFactSkip(num int) int {
    if num % 2 == 0 {
        return primeFact(num/2)
    }
    smallestDivisor := 3
    for num % smallestDivisor != 0{
        smallestDivisor += 2
    }
    if smallestDivisor == num {
        return smallestDivisor
    }
    return primeFact(num/smallestDivisor)
}
