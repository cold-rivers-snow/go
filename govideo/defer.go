package main

import "fmt"

func main(){
    //defer 延迟
    defer fmt.Printf("world")
    fmt.Printf("hello")

    for a := 0;a < 10;a++{
        defer fmt.Printf("a = %d\n",a)
    }

    defer fmt.Printf("2")
    defer fmt.Printf("1")

}