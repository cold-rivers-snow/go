package main

import  "fmt"

//src 源码  
//pkg 存放go install命令构建安装后的代码包的归档文件
//bin 存放go install命令之后的可执行文件
//在/etc/profile的GOPATH 后添加工作区 source 生效
//标识符 字母 数字 下划线  且必须字母开头
//每一个标识符使用前必须声明 使用其他包中的变量之间必须用.隔开 os.O_RDONLY。
//使用另一个包，必须导入，使用其标识符必须是可导入的。
//可导出标识符：第一个字符必须大写，标识符必须被声明在代码包中的变量或类型名称，或属于某个结构体类型的字段名称或方法名称。
//标识符第一个字母大写，说明其访问权限公开，小写访问权限私有，这是go通过标识符第一个字母来判断的。
//预定义标识符：所有的基本类型，接口类型error，常量true 、false和iota，所有的内建函数：append、cap、close、complex、copy、delete、imag、len、make、new、panic、print、println、real、recover。
//空标识符用一个_表示。主要用于某个函数返回值或变量不需要使用时 或 某个包的标识符不需要对其操作时用空去接，避免了编译出错。 import _ "runtime/cgo" 这个包代表一个标准库代码包的标识符。
// 关键字：
//程序声明：import，package
//程序实体声明和定义：chan，const，func，interface，map，struct，type，var
//程序流程控制：go，select，break，case，continue，default，defer，else，fallthrough，for，goto，if，range，return，switch
//基本数据类型：string，bool，byte，rune，int/uint，int8/uint8,int16/uint16,int32/uint32，int64/uint64,float32,float64,complex64,complex128.
//复杂数据类型：Array，Struct，Function，Interface，Slice，Map，Channel，Pointer。
//动态类型：运行时与该变量绑定在一起的值的实际类型。Interface可以是动态的。  
//一个类型的潜在类型具有传递性：type Mystring string; type iString Mystring; 说明潜在的iString也是string类型。
//操作符：||   &&   ==   !=  <  <=  >  >=  +  -  |  ^  *  /  %  <<  >>  &  &^(按位清除 二元操作符 算术)  !  <-(接收操作 一元操作符 接收操作符  通道类型)  ! <- 也可以用于二元操作符。
//&^清除操作符，根据第二个操作数的二进制对第一个操作数的二进制进行相应的清零操作。如果第二个操作数的某个二进制位上的数值为1,那么第一个操作数的对应位二进制为0,否则，不变。不会改变第一个操作数的值。
//如：00000111 &^ 00001101 = 00000010 (7 &^ 13 = 2)  ;00001101 &^ 00000111 = 00001000(13 &^ 7 = 8)
//^ 作为一元操作符时，与其整数类型的最大值进行异或，如：^uint8(1) = 254   (00000001 ^ 11111111 = 11111110); ^1 = 2 (00000001 ^ 11111111 = 11111110)
//空值（nil）
//接收操作符 <-   对于通道类型的值ch，表达式<-ch含义是从此通道中接收一个值。前提是通道的方向必须允许接收操作，并且该操作的结果的类型必须与通道元素的类型之间存在可赋予的关系。这个表达式回被阻塞直到通道中有一个值可用。
//从一个通道类型的空值（即nil）接收值的表达式将会永远被阻塞。
//从一个被关闭的通道类型值接收值会永远成功并立即返回一个其元素类型的零值。

func main(){
	fmt.Println("Hello world")
}