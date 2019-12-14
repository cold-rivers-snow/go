//const b string = "abc"    //显式类型定义
//const b = "abc"           //隐式类型定义
//多个相同类型声明简写： const c_name1, c_name2 = value1, value2

package main

import "fmt"

func main(){
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"//多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d",area)
	println()
	println(a,b,c)
}

//常量用于枚举
//const (
//	Unknown = 0
//	Female = 1
//	Male = 2
//)
