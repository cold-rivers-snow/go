package main

import  "fmt"

//src 源码  
//pkg 存放go install命令构建安装后的代码包的归档文件
//bin 存放go install命令之后的可执行文件
//在/etc/profile的GOPATH 后添加工作区 source 生效
//标识符 字母 数字 下划线  且必须字母开头
//每一个标识符使用前必须声明 使用其他包中的变量之间必须用.隔开 os.O_RDONLY。


func main(){
	fmt.Println("Hello world")
}