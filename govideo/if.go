package mai
import "fmt"

func main(){
    //定义局部变量
    var a int = 10

    //使用if语句判断布尔表达式
    if a < 20{
        //如果条件为true则执行以下语句
        fmt.Printf("a 小于20\n")
    }
    fmt.Printf("a 的值为：%d\n",a)
}