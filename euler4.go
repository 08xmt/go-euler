package main

import (
    "fmt"
    "time"
    "strconv"
)
//The prime factors of 13195 are 5, 7, 13 and 29..
//What is the largest prime factor of the number 600851475143 ?.
func main() {
    start := time.Now()
    res := largestPalindrome(999)
    fmt.Println(time.Since(start))
    fmt.Println(res)
}

func largestPalindrome(maxNum int) int {
    largest := 0
    j := maxNum
    for i := maxNum; i > 0; i-- {
        j = i
        if largest > i*j {
            return largest
        }
        for j = i; j > 0; j--{
            if isPalindrome(i*j){
                if largest < i*j {
                    largest = i*j
                }
                break
            }
        }

    }
    return largest
}

//TODO: Implement function that calculates palindromes without using string converts
func isPalindrome(num int) bool {
    numStr := strconv.Itoa(num)
    strLen := len(numStr)
    if strLen % 2 != 0 {
        return false
    }
    for i := 0; i < strLen/2; i++ {
        if numStr[strLen/2+i] != numStr[strLen/2-i-1]{
            return false
        }
    }
    return true
}


