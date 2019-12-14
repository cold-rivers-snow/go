package main

import "fmt"

//二维数组初始化
var a = [3][4]int {
	{0, 1, 2, 3},	//第一行索引为0
	{4, 5, 6, 7},	//第二行索引为1
	{8, 9, 10, 11},	//第三行索引为2
}

//访问二维数组
//var val := a[2][3]
var value int = a[2][3]

func main(){
	//数组-5行2列
	var a = [5][2]int {{0,0},{1,2},{2,4},{3,6},{4,8}}
	var i, j int

	//输出数组元素
	for i = 0;i < 5;i++{
		for j = 0;j < 2;j++{
			fmt.Printf("a[%d][%d] = %d\n",i, j, a[i][j])
		}
	}
}