package main

import "fmt"

func main(){
    var a int = 21
    var c int

    c = a
    fmt.Printf("第1行 - = 运算符实例，c值=%d\n",c)

    c += a
    fmt.Printf("第2行 - += 运算符实例，c值=%d\n",c)

    c -= a
    fmt.Printf("第3行 - -= 运算符实例，c值=%d\n",c)

    c *= a
    fmt.Printf("第4行 - *= 运算符实例，c值=%d\n",c)

    c /= a
    fmt.Printf("第5行 - /= 运算符实例，c值=%d\n",c)

    c = 200;

    c <<= 2
    fmt.Printf("第6行 - <<= 运算符实例，c值=%d\n",c)
    
    c >>= 2
    fmt.Printf("第7行 - >>= 运算符实例，c值=%d\n",c)

    c &= 2
    fmt.Printf("第8行 - &= 运算符实例，c值=%d\n",c)

    c ^= 2
    fmt.Printf("第9行 - ^= 运算符实例，c值=%d\n",c)

    c |= 2
    fmt.Printf("第10行 - |= 运算符实例，c值=%d\n",c)


}