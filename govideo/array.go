package main

import "fmt"
//数组初始化
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//忽略[]中的数字不设置数组大小
var balance1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

func main(){
	var n [10] int //n式一个长度为10的数组
	var i, j int
	

	//数组n初始化元素
	for i = 0;i < 10;i++{
		n[i] = i + 100;
	}

	//输出每个数组元素的值
	for j = 0;j < 10;j++{
		fmt.Printf("Element[%d] = %d\n",j, n[j])
	}
}