package main

import "fmt"

func main(){
	//下面两个都是切片
	array := [...]int{10,20,30,0,0}
	slice := array[1:3:5] //[low（下标起点）,high（下标终点，但不包括）,容量]（左闭右开）
	//len = high - low; cap = max - low
	fmt.Println(cap(slice),len(slice))
	fmt.Println(slice)
}
