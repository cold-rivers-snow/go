package main

import "fmt"

func main(){
    //定义局部变量
    var a int = 10

    Loop: for a < 20{
        if a == 15{
            //跳出迭代
            a = a + 1
            goto Loop
        }
        fmt.Printf("a的值为：%d\n",a)
        a++
    }
}