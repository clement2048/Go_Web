package main
import "fmt"

//定义一个结构体类型
type Student struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}

func main(){
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1,"mike",'m',6,"home"}
	fmt.Println("s1 = ",s1)
}