package main

import (
	"fmt"
	"time"
)

func longTack(signal chan int) {
	for {
		fmt.Println("longtask is running")
		v := <-signal

		if v == 1 {
			break
		}else{
			t := v
			signal <- t-1
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println("longtask is finish")
}

func main() {
	sig := make(chan int)
	go longTack(sig)

	time.Sleep(3 * time.Second)

	sig <- 1
	time.Sleep(1 * time.Second) //防止主进程跑太快，协程不能输出
}
