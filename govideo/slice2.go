package main

import "fmt"

func main(){
	//创建切片
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	//打印原始切片
	fmt.Println("numbers ==",numbers)

	//打印子切片从索引1（包含）到索引4（不包含）
	fmt.Println("numbers[1:4] ==",numbers[1:4])

	//默认下限为0
	fmt.Println("numbers[:3] ==",numbers[:3])

	//默认上限为len（s）
	fmt.Println("numbers[4:] ==",numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	//打印子切片从索引0（包含）到索引2（不包含）
	numbers2 := numbers[:2]
	printSlice(numbers2)

	//打印子切片从索引2（包含）到索引5（不包含）
	numbers3 := numbers[2:5]
	printSlice(numbers3)

	/*
	切片的capacity初始值取决于切片的初始化方式：

1.通过make函数初始化一个切片时，capacity由我们自己定义。

2.通过字面量初始化一个切片时，capacity默认等于该切片的长度。

3.对数组或切片执行array[start:end]操作生成切片时，切片的capacity总等于源数组/源切片的capacity减去start的值，比如array原本的capacity为4，s0 := array[1:],则s0的capacity为3。

	*/

}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n",len(x),cap(x),x)
}