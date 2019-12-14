package main

import "fmt"

func main(){
    //定义局部变量
    var a int = 10

    for a < 20 {
        if a == 15{
            //跳出此次循环
            a = a + 1
            continue;
        }
        fmt.Printf("a的值为：%d\n",a)
        a++;
    }
}