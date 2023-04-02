package main

import(
	"fmt"
)

func add(a, b int) int{
	return a + b
}

func minus(a, b int) int{
	return a - b
}

// 函数也是一种数据类型， 通过type给一个函数类型起名

//没有函数名字和大括号
//Functype是一种函数类型
type Functype func(int, int) int

func main(){
	var result int
	result = add(1,1) // 传统调用方法
	fmt.Println("result = ",result)	
	var ftest Functype
	ftest = add
	result = ftest(20,10)
	fmt.Println("result = ", result) 
}