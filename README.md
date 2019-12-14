此项目是go语言学在菜鸟教程中的一些笔记，适用于一些初学者，以后go相关的学习和项目将持续更新。

[go语言](http://www.runoob.com/go/go-tutorial.html)

go的IDE 

LiteIDE：

LiteIDE 是一款开源、跨平台的轻量级 Go 语言集成开发环境（IDE）。
支持的 操作系统

    Windows x86 (32-bit or 64-bit)
    Linux x86 (32-bit or 64-bit)

[下载地址](http://sourceforge.net/projects/liteide/files/)

[源码地址](https://github.com/visualfc/liteide)

Eclipse :
    首先下载并安装好 Eclipse
    [下载goclipse插件](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#installation) 
    下载 gocode，用于 go 的代码补全提示
    [gocode的github地址](https://github.com/nsf/gocode)

    

    在 Windows下要安装 git，通常用 msysgit。

    再在 cmd 下安装：

    go get -u github.com/nsf/gocode

    也可以下载代码，直接用 go build 来编译，会生成 gocode.exe
    下载 MinGW 并按要求装好
    配置插件

    Windows->Reference->Go

    (1)、配置 Go 的编译器

    1.4.eclipse2

    设置 Go 的一些基础信息

    (2)、配置 Gocode（可选，代码补全），设置 Gocode 路径为之前生成的 gocode.exe 文件

    1.4.eclipse3

    设置 gocode 信息

    (3)、配置 GDB（可选，做调试用），设置 GDB 路径为 MingW 安装目录下的 gdb.exe 文件

    1.4.eclipse4

    设置 GDB 信息

    测试是否成功

    新建一个 go 工程，再建立一个 hello.go。如下图：

    1.4.eclipse5

    新建项目编辑文件

    调试如下（要在 console 中用输入命令来调试）：

    1.4.eclipse6

    go程序设计语言

