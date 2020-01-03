package set

import (
	"bytes"
	"fmt"
)
//go中没有Set集合数据类型，有Map的字典数据类型。//Map使用哈希表实现的。
//Set和Map 相同点:
//1、它们的元素都是不可重复的
//2、它们都是只用迭代的方式取出其中的所有元素
//3、对它们中的元素进行迭代的顺序都是与元素插入顺序无关的，同时也不保证任何有序性。
//区别：
//1、Set的元素是一个单一值，而Map的元素是一对键值
//2、Set的元素不可重复指的是不能存在任意两个单一值相等的情况，Map的元素不可重复指的是任意两个键值对中的键的值不能相等


type HashSet struct{
	m map[interface{}bool		
}



//HashSet 类型的元素可以是任何类型的，所以，m的键类型为interface{}，用m的键来存储HashSet的元素值。因此，选用值占用空间最小的类型来作为m的值的元素类型。
//使用bool的好处：
//1、bool占用空间小，1字节
//2、bool的值只有两个。且这两个都是预定义的常量
//3、利于判断字典类型值中是否存在某个键。
//可以使用v,ok := m["a"]		来判断

//由于map的空值为nil，因此，如果new(HashSet).m也为nil，
//创建和初始化的函数
fucn NewHashSet() *HashSet{
	return &HashSet{m: make(map[interface{}]bool)}
}
//使用*HashSet是为了在方法集合包中包含调用接收者类型为HashSet或*HashSet的所有方法。

//基本功能
//添加元素值
func (set *HashSet) Add (e interface{}) bool{
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}		//接收者类型为*HashSet	，定义为指针，减少复制接收者对系统能够资源的耗费。



//删除元素值
func (set *HashSet) Remve(e interface{}) {
	delete (set.m,e)
}
//清除所有元素值
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}//go中没有对Map的清除操作，如果利用循环的删除，这样做，并发访问和修改的情况下会引发问题，因此，采用一个巧妙的方法，重新初始化。

//判断是否包含某个元素值
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}
//字典类型的键值不能是函数类型、字段类型或切片类型。其interface{}内部会先动态的转换成实际的类型值，在与之相对应的hash对该值进行计算，因此，总能正确的计算除hash值。

//获取元素值的数量
func (set *HashSet) Len() int {
	return len(set.m)
}

//判断与其他HashSet类型值是否相同
//包含元素值应完全相同。
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//获取所有元素值，即生成可迭代的快照
//快照：目标值某一时刻的影像。
//从元素可迭代且顺序可确定的数据类型值中选取一个，以单值作为元素，因此不能为字典类型。又由于HashSet的数量不确定，数组排除，
//因此，快照可以是切片或通道类型。
//以切片类型为例
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{},initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		}else {
			snapshot = append(snapshot,key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot  = snapshot[:actualLen]		//初始化后，没有的元素为nil，如果快照尾部有若干个没意义的值则去掉。
	}
	return snapshot
}
//获取自身的字符串表示形式
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v",key))
	}
	buf.WriteString("}")
	return buf.String()
}

//根据集合代数中的描述，如果集合A真包含了集合B，那么就可以说集合A是集合B的一个超集。因此这个方法名称可以确定为IsSuperset
func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen == 0 || otherLen == 0 {
		return false
	} 
	if oneLen > 0 && otherLen == 0{
		return true
	}
	for _,v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//以下几个只是个人的简介，不一定正确
//并集运算指两个集合中所有的元素合并起来并组成一个集合。Union
func (set *HashSet) Union(other *HashSet) *HashSet {
	if other == nil  || set == nil {
		return set
	}
	oneLen := set.Len()
	otherLen := other.Len()
	for _,v := range other.Elements() {
		if !set.Contains(v) {
			continue
		} else {
			set.append(set,v)
		}
	}
	return set
}
//交集运算指找到两个集合中共有的元素并把它们组成一个集合。Intersect
func (set *HashSet) Intersect(other *HashSet) *HashSet {
	if other == nil {
		return nil
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen < otherLen {
		sv := make([]interface{},oneLen)
	} else {
		sv := make([]interface{},otherLen)
	}
	for _,v := range other.Elements() {
		if set.Contains(v) {
			sv[v] = true
		}
	}
	return sv
}
//集合A 对集合B进行差集运算指找到只存在于A中但不存在于集合B中的元素并把它们组成一个集合。Difference 
func (set *HashSet) Difference(other *HashSet) *HashSet {
	if other == nil {
		return nil
	}
	sv := set.Union(other)
	oneLen := set.Len()
	ss := make([]interface{},oneLen)
	for _,v := range sv.Elements() {
		if set.Contains(v) {
			continue
		} else {
			ss[v] = true
		}
	}
	return ss
}
//对称差集与差集区别：对称差集指找到只存在于集合A 中但不存在于集合B中的元素，再找只存在B中不存在A中的元素，最后把它们合并成一个集合。SymmetricDifference
func (set *HashSet) SymmetricDifference(other *HashSet) *HashSet {
	if other == nil {
		return nil
	}
	sv := set.Union(other)
	svo := other.Union(set)
	oneLen := set.Len()
	otherLen := other.Len()
	ss := make([]interface{},oneLen + otherLen)
	for _,v := range set.Elements() {
		if sv.Contains(v) {
			continue
		} else {
			ss[v] = true
		}
	}
	for _,v := range other.Elements() {
		if sv.Contains(v) {
			continue
		} else {
			ss[v] = true
		}
	}
	return ss
}
//满足条件：
//1、它们都接受一个名为other且类型为*HaseSet的参数值。
//2、在方法中不得修改接收者Set的值和参数other的值。
//3、它们的结果类型都应为*HashSet。
//4、尽可能地利用已有的附属于*HashSet类型的方法。

//进一步重构
type Set interface{
	Add(e interface{}) bool
	Remve(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func IsSuperset(one Set,other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := otherLen.Len()
	if otherLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _,v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}



func main(){

}