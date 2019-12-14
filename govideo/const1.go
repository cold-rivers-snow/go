package main

//常量表达式中，必须是内置函数
import "unsafe"

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
	println(a, b, c)
}

//iota关键字，为常量0,const每增加一行声明iota计数一次

//const(
//	a = iota
//	b = iota
//	c = iota
//)

//简写为
//const(
//	a = iota
//	b 
//	c 
//)
