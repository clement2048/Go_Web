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

// ����Ҳ��һ���������ͣ� ͨ��type��һ��������������

//û�к������ֺʹ�����
//Functype��һ�ֺ�������
type Functype func(int, int) int

func main(){
	var result int
	result = add(1,1) // ��ͳ���÷���
	fmt.Println("result = ",result)	
	var ftest Functype
	ftest = add
	result = ftest(20,10)
	fmt.Println("result = ", result) 
}