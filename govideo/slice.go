//切片：“动态数组”
//定义：var id []type  或者使用make函数
//var slice1 []type = make([]type,len)
//简写：slice1 := make([]type,len)
//make([]T, length, capacity)

//切片初始化
//s := []int{1,2,3}
/*
	s := arr[:]
	s := arr[startIndex:endIndex]
	s := arr[startIndex:]	//从startIndex到最后一个元素
	s := arr[:endIndex] 	//从开始到endIndex
	s1 := s[startIndex:endIndex]
	s := make([]int,len,cap)
	切片可以索引，并可以使用len（）方法获取长度
	使用cap()测量切片最长可到达多少
*/

package main

import "fmt"

func main(){
	var numbers = make([]int,3,5)

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n",len(x),cap(x),x)
}
