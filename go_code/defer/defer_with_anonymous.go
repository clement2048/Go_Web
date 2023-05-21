package main
import "fmt"

func main(){
	a := 10
	b := 20
	defer func (a,b int)  {
		fmt.Println("a = ", a," b = ",b)
	}(a,b) //把参数传递过去，已经先传递参数，只是没有调用，这里是值传递

	a = 111
	b = 222
	fmt.Printf("%d,%d\n",a,b)
}

func main01(){
	a := 10
	b := 20
	defer func ()  {
		fmt.Println("a = ", a," b = ",b)
	}()

	a = 111
	b = 222
	fmt.Printf("%d,%d\n",a,b)
}