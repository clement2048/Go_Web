package main

import(
	"fmt"
)

func main(){
	//在定义是直接调用，只能调用一次
	a := 10
	str := "mike"
	//匿名函数，没有函数名字
	//:= 自动推导类型
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	//调用函数
	f1()

	//给一个函数类型起别名
	type FuncType func () //函数没有参数，没有返回值
	var f2 FuncType 
	f2 = f1
	f2()

	//定义匿名函数同时调用
	func(){
		fmt.Printf("a = %d, str = %s\n",a,str)
	}() //后面的()代表调用此匿名函数

	//带参数的匿名函数
	f3 := func (i, j int)  {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(1,2)
	//定义的同时调用函数
	func (i, j int)  {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(10,20) 
	
	//匿名函数，有参有返回值
	x, y := func(i, j int) (max, min int){
		if i>j {
			max = i
			min = j
		}else{
			min = i
			max = j
		}
		return 
	}(10, 20)
	fmt.Println(x,y)
}