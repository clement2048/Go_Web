package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

// 这是一个func函数啥的

// 这里测试一下
// 这里也测试一下
type Functype func(int, int) int

func main() {
	var result int
	result = add(1, 1) // 这里还在测试
	fmt.Println("result = ", result)
	var ftest Functype
	ftest = add
	result = ftest(20, 10)
	fmt.Println("result = ", result)
}
