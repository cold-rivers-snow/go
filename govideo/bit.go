package main

import "fmt"

func main(){
    var a uint = 60
    var b uint = 13
    var c uint = 0

    c = a & b
    fmt.Printf("第一行 - c的值为%d\n",c)

    c = a | b
    fmt.Printf("第二行 - c的值为%d\n",c)

    c = a ^ b
    fmt.Printf("第三行 - c的值为%d\n",c)

    c = a << 2
    fmt.Printf("第四行 - c的值为%d\n",c)

    c = a >> 2
    fmt.Printf("第五行 - c的值为%d\n",c)
}