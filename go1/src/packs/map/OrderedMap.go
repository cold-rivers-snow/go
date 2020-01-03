package orderedmap

import (
	"fmt"
	"reflect"
	"sort"
)

type OrderedMap struct {
	keys []interface{}
	m map[interface{}] interface{}
}


func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}

//比较两个元素值的大小并返回判断结果
//只有当值具有有序性才可与其他同类型比较。
func (omap *OrderedMap) Less(i,j int) bool {
	
}

func (omap *OrderedMap) Swap(i,j int) {
	omap.keys[i], omap.keys[j] = omap.keys[j],omap.keys[i]
}

//keys有序，方便比较。
//keys不是具体类型，运行时确定
//对keys进行添加、删除、获取等操作。
//keys应固定顺序获取
//keys自动排序。
//keys有序后，确定某一个元素位置
//运行时确定keys的元素类型，则也可获知这个元素类型。
//运行时获取到被用于比较keys值中不同元素的大小具体方法。
//因此定义接口：
type Keys interface{
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int,contains bool)
	ElemType() reflect.Type
	CompareFunc() func(interface{},interface{}) int8
}

type myKeys struct {
	container []interface{}
	compareFunc func(interface{},interface{}) int8
	elemType reflect.Type
}

int64Keys := &myKeys{
	container : make([]interface{},0),
	compareFunc : func(e1 interface{},e2 interface{}) int8 {
		k1 := e1.(int64)
		k2 := e2.(int64)
		if k1 < k2 {
			return -1
		}else if k1 > k2 {
			return 1
		}else {
			return 0
		}
	},
	elemType : reflect.TypeOf(int64(1))
}

func (keys *myKeys) Len() int {
	return len(keys.container)
}

//比较两个元素值的大小并返回判断结果
//只有当值具有有序性才可与其他同类型比较。
func (keys *myKeys) Less(i,j int) bool {
	return keys.compareFunc(keys.container[i],keys.container[j]) == -1
}

func (keys *myKeys) Swap(i,j int) {
	keys.container[i], keys.container[j] = keys.container[j],keys.container[i]
}

//判断是否可接收
func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != keys.elemType {
		return false
	}
	return true
}

func Sort(data Interface) {

}

func (keys *myKeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container,k)
	sort.Sort(keys)
	return true
}

func Search(n int,f func(int) bool) int{

}

//Search的func参数函数
func(i int) bool {
	return keys.compareFunc(keys.container[i],k) >= 0
}

//Search之后进行判断
if index < keys.Len && keys.container[index] == k {
	contains = true
}

func (keys *myKeys) Remove(k interface{}) bool {
	if !keys.Search(k) {
		return false
	}else {
		keys.container = append(keys.container[0:k],keys.container[k+1]...)
	}
	return true
}//.......
func (keys *myKeys) Clear() {
	keys.container = make([]interface{},0)
}

func (key *myKeys) Get(index int) interface{} {
	if index >= keys.Len() {
		return nil
	}
	return keys.container[index]
}

func (key *myKeys) GetAll() []interface{} {
	initialLen := len(keys.container)
	snapshot := make([]interface{},initialLen)
	actualLen := 0
	for _,key := range keys.container {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		}else {
			snapshot = append(snapshot,key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (keys *myKeys) ElemType() reflect.Type{
	return key.elemType
}

func (keys *myKeys) CompareFunc() CompareFunction {
	return keys.compareFunc
}

func NewKeys (
	compareFunc func(interface{},interface{}) int8,
	elemType reflect.Type) keys {
	return &myKeys {
		container: make([]interface{},0),
		compareFunc: compareFunc,
		elemType:elemType,
	}	
	}

	type myOrderedMap struct {
		keys Keys
		elemType reflect.Type
		m map[interface{}]interface{}
	}

	type OrderedMap interface{
		Get(key interface{}) interface{}
		Put(key interface{},elem interface{}) (interface{},bool)
		Remove(key interface{}) interface{}
		Clear()
		Len() int
		Contains(key interface{}) bool
		FirstKey() interface{}
		LastKey() interface{}
		HeadMap(toKey interface{}) OrderedMap
		SubMap(fromKey interface{},toKey interface{}) OrderedMap
		TailMap(fromKey interface{}) OrderedMap
		Keys() []interface{}
		Elems() []interface{}
		ToMap() map[interface{}]interface{}
		KeyType() reflect.Type
		ElemType() reflect.Type 
	}

	

