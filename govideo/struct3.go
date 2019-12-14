package main

import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	var Book1 Books		//声明Book1为Books类型
	var Book2 Books 	//声明Book2为Books类型

	//book1描述
	Book1.title = "Go语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go语言教程"
	Book1.book_id = 6495407

	//book2描述
	Book2.title = "Python类型"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python语言教程"
	Book2.book_id = 6495700

	//打印Book1信息
	printBook(&Book1)

	//打印Book2信息
	printBook(&Book2)

}

func printBook(book *Books){
	//打印Book1信息
	fmt.Printf("Book  title : %s\n",book.title)
	fmt.Printf("Book  author: %s\n",book.author)
	fmt.Printf("Book  subject:%s\n",book.subject)
	fmt.Printf("Book  book_id：%d\n",book.book_id)
}