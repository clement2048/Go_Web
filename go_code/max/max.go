package main

import "fmt"


func max(x1, x2 int) int {
	var result int
	if x1 > x2 {
		result = x1
	} else {
		result = x2
	}
	return result
}

func swap (x,y int) (int,int){
	return y,x
}


func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("the max number is : %d \n", ret)
	swap(a,b)
	fmt.Printf("%d,%d",b,a)
}