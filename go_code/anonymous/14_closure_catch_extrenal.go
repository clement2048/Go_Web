package main

import "fmt"

func main(){
	a := 10
	str := "mike"

	func(){
		//闭包是以引用方式捕获外部变量
		a = 66
		str = "go"
		fmt.Println("内部:a = ", a, " str = ",str);
	}()
	fmt.Println("外部:a = ", a, " str = ", str)
}