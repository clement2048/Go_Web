package main

import(
	"fmt"
)

//回调函数，函数有一个参数是函数类型
//计算器，可以进行四则运算
//主要就是多态，调用同一个接口，但是不同的表现
//先有想法，后面再实现功能
type Functype func(int, int) int

func add (a,b int) int{
	return a+b
}
func Calc(a, b int, fTest Functype) int {
	fmt.Println("Calc") //可以解耦，实现利用多个函数
	r := fTest(a,b)
	return r
}

func main(){
	a := Calc(1,2,add)
	fmt.Print(a)
}