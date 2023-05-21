package main

import "fmt"
//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func()int {
	var x int
	return func()int{
		x++
		return x*x
	}
}

func main(){
	//返回值为一个匿名函数，所以返回一个函数类型,f会调用返回的匿名函数
	//他不关心那些捕获了的变量和常量是否已经超出了作用域
	//所以只有闭包还在使用它，这些变量就还会存在
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
}
func test01() int{
	var x int //没有初始化，值为0
	x++
	return x * x
}

func tmain(){
	fmt.Println(test01())
}
