/*
//类型相同多个变量，非全局变量
var vname1, vname2, vname3  type
vname1, vname2, vname3 = v1, v2, v3

var vname4, vname5, vname6 = v4, v5, v6

vname7, vname8, vname9 := v7, v8, v9

//这种因式分解关键字的写法一般用于声明全局变量
var（
	vname1 v_type1
	vname2 v_type2
）
*/
package main

var x, y int
var (
    //这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
//定义必须被使用



func main(){
	//交换两个变量的值,类型必须相同，在函数体内
	c, d = d, c
	_, d = 3, 4	//_为被抛弃值，是一个只写变量，不能得到_的值
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
