package main

import (
	"fmt"
	"sort"
)
type Interface interface{
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}//接口声明，这个接口包含3个方法
//一个接口类型嵌入到另一个接口类型
type Sortable interface{
	sort.Interface
	Sort()
}//此接口实质上声明了4个方法，但是接口不能嵌入自身，无论是直接嵌入还是间接嵌入。嵌入的方法不能重名。
// interface{}//空接口

type SortableStrings [3]string
//type SortableStrings []string


func (self SortableStrings) Len() int {
	return len(self)
}
func (self SortableStrings) Less(i,j int) bool{
	return self[i] < self[j]
}
func (self SortableStrings) Swap(i,j int) {
	self[i],self[j] = self[j],self[i]
}
//有了方法实现，就可以使用sort.Interface接口了。
// _ , ok := interface{}(SortableStrings{}).(sort.Interface)	//需要导入包sort，验证接口类型实现的直接方法。
//一个接口类型可以被任意数量的数据类型实现，相反，一个数据类型也可以同时实现多个接口类型。
func (self SortableStrings) Sort(){
	sort.Sort(self)
}

func main(){
	
	_ , ok := interface{}(SortableStrings{}).(sort.Interface)
	_ , ok2 := interface{}(SortableStrings{}).(Sortable)
	//如果把SortableStrings改为一个指针类型，则测试改为：_,ok2 := interface{}(&SortableStrings{}).(Sortable)
	if(ok){
		fmt.Println("ok")
	}
		
	if(ok2){
		fmt.Println("ok2")
	}
		
	ss := SortableStrings{"2","3","1"}
	ss.Sort()
	fmt.Printf("Sortable Strings:%v\n",ss)

	ints := []int{1,2,3,4,5}

	for i,d := range ints{
		fmt.Printf("%d:%d\n",i,d)
	}
}