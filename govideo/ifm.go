package main

import "fmt"

func main(){
    //定义局部变量
    var a int = 100
    var b int = 200

    //判断条件
    if a == 100 {
        //if语句为true
        if b == 200 {
            //if语句为true
            fmt.Printf("a 的值为100,b的值为200\n")
        }
    }
    fmt.Printf("a值为：%d\n",a)
    fmt.Printf("b值为：%d\n",b)
}