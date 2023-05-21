package main
import(
	"fmt"
	"time"
)

func add

func initchannel(n int,c1 chan int){
	for i:= 2; i <n; i++{
		c1 <- i
	}
	c1.close()
}

func main(){
	n := 100
	c1 = make(chan int)
	go initchannel(n,c1)

	go
	for i:= 2; i <n; i++{
		go 
		c2 = make(chan int)
	}
}