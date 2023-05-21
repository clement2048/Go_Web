package main

import (
    "fmt"
)

func generate(nums chan<- int) {
    for i := 2; ; i++ {
        nums <- i
    }
}

func filter(nums <-chan int, primes chan<- int, prime int) {
    for {
        num := <-nums
        if num%prime != 0 {
            primes <- num
        }
    }
}

func main() {
	n := 20
    nums := make(chan int)
    go generate(nums)

    for i := 0; i < n; i++ {
        prime := <-nums
        fmt.Println(prime)
        primes := make(chan int)
        go filter(nums, primes, prime)
        nums = primes
    }
}
