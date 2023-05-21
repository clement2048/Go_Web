package main

import(
	"fmt"
)

func Generator(ch chan<- int){
	for i := 2; ; i++{
		ch <- i
	}
}

func Filter(ch chan<- int, ch1 chan<- int, prime int){
	for{
		num := <-ch
		if(num% prime != 0){
			ch1 <- num
		}
	}
}

func main1(){
	n := 20
	ch := make(chan int)
	go Generator(ch)
	for i := 0; i < n; i++{
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
	