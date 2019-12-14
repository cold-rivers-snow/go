package main

import "fmt"

//定义结构体
type Circle struct{
	radius float64
}

func main(){
	var c1 Circle
	c1.radius = 10.0
	fmt.Println("圆的面积=",c1.getArea())
}

//该method属于Circle类型对象中的方法
func (c Circle) getArea() float64{
	//c.radius即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}