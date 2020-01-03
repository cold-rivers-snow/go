package main

import  (
	"fmt"
	"testing"
)
func main(){
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
//[]int{1,2,3,4,5} 切片 []int{1,2,3,4,5}[2] 结果为3,取出索引为2的元素。这种表达式称为索引表达式
//uint8 类型的v1，与int类型的v2,相加 int(v1)+v2
//内建函数len   len(v2)
//[]int{1,2,3,4,5}取出第二到第四个元素[]int{1,2,3,4,5}[1:4]
//判断int8类型的num是否为int类型：interface{}(num).(int).
//os.Open("/etc/profile")
//如果a不是字典类型，a[x]中x必须是int类型或整数字面量，且大于等于0小于其长度。
//切片a[x:y:z]  x：切片元素下界索引，y：切片元素上界索引，z：容量上界索引。
//切片表达式的作用是截取字符串。a[x:y] 取的长度为y-x，左闭右开。a[:x]截取x之前的所有元素。a[x:]截取x之后的所有元素。a[:]a中所有元素。
//go语言采用utf8编码的。utf8的一个中文字符包含3个字节。 
//断言:一个接口类型x和一个类型T 对应的类型断言：x.(T) 其作用判断x不为nil且存储在其中的值是T类型的。
//interface{} 空接口。
//v,ok := x.(T)   其中v为求值结果，ok为判断断言是否成功。 
//...T 可变参数类型声明。func appendIfAbsent(s []string,t ...string) []string
//调用appendIfAbsent([]string{"A","B","C"},"C","B","E")
//appendIfAbsent([]string{"A","B","C"},[]string{"C","B","E"}...)

//数据类型
// 名称				字节大小			零值				说明
// bool					1							false 			布尔类型
// byte					1							0					字节类型（8位二进制代表的无符号整数） 可以看作uint8类型的别名
// rune					4							0					rune类型 专门用于存储Unicode编码（32位二进制代表的有符号整数）可以看作uint32类型的别名  存储的Unicode代码点以U+作为前缀
// int/uint        	  -								0					32位下4字节，64位下8字节
// int8/uint8   	  1							0
// int16/uint16    2							0
// int32/uint32    4							0
// int64/uint64     8							0
// float32  				4						0.0
// float64					8						0.0
// complex64			8						0.0+0.0i		实部和虚部分别用float32类型表示
// complex128		8						0.0+0.0i		实部和虚部分别用float64类型表示
// 	string				-						""

//数组    [n]T   长度不同，类型相同，两个数组，依然不同。
// [6]string{"Go","Python","Java","C","C++","PHP"}		//一个长度为6,元素类型为string类型，且包含6个元素值。
// [6]string{0:"Go",1:"Python",2:"Java",3:"C",4:"C++",5:"PHP"}	//顺序也可以打乱
// [6]string{1:"Python",3:"C",0:"Go",4:"C++",2:"Java",5:"PHP"}//也可以显示一部分
// [6]string{4:"Python",5:"C",0:"Go","C++","Java",3:"PHP"}
//满足条件：指定的索引值在该数组的有效范围内，同时索引是递增的，[6]string{"Go","Python","Java","C",5:"C++","PHP"}	非法的
//指定索引值不能与其他元素值的索引值重复[6]string{0:"Go","Python",1:"Java","C","C++","PHP"}	
// [...]string{"Go","Python","Java","C","C++","PHP"}	//不显式指定长度
// len([...]string{"Go","Python","Java","C","C++","PHP"})//获取长度
// [...]string{"Go","Python","Java","C","C++","PHP"}[0] //为"GO"
// array1 := [6]string{"Go","Python","Java","C","C++","PHP"}
// array1[2] = "Clojure" //则[6]string{"Go","Python","Clojure","C","C++","PHP"}
//切片类型 []T 
// []string{"Go","Python","Java","C","C++","PHP"}	
// []string{4:"Python",5:"C",0:"Go","C++","Java",3:"PHP"}//与数组不同的是，只需要索引值不重复即可，没有长度限制。
// //切片的零值为nil
// len([]string{"Go","Python","Java","C","C++","PHP"}	)//切片长度 长度为6
// len([]string{8:"Go",2:"Python","Java","C","C++","PHP"} //长度为9
//切片的底层实现是总持有一个对某数组值的引用。对切片值中元素值的修改，实质上是对底层数组上对应元素的修改。
//切片值类似于指向底层数组的指针。切片值的容量与底层数组的长度有关。用cap获取容量。cap([]string{8:"Go",2:"Python","Java","C","C++","PHP"}	)
//一个切片底层的数据结构包含一个指向底层数组的指针类型值、一个代表切片长度的int类型和一个代表切片容量的int类型值。
//*Elem ptr , int len , int cap
//使用复合字面量初始化的一个切片值类型时，首先创建这个切片值的底层数组，这个底层数组与这个切片值有着相同的元素类型、元素值及其排列顺序和长度。
array1 := [...]string{"Go","Python","Java","C","C++","PHP"}	
slice1 := array1[:4]			//生成一段新的切片，其中slice1的底层数组就是array1。切片表达式的作用是创建一个新的切片值。
//slice1的切片其中*Elem 指针指向array1的首部地址 ,  len  = 4,  cap = 6
slice2 := array1[3:]//slice2的切片其中*Elem 指针指向array1的第3个元素的地址 ,  len  = 3,  cap = 3,。因此，一个切片值的容量并不一定就是其底层数组的长度。
//slice1[4]	//编译错误，超过了其长度
//slice1 = slice1[:cap(slice1)]//扩展窗口，扩展其底层数组的容量大小，一个切片的窗口只能向一个方向扩展。即索引值递增的方向。切片值不允许负整数字面量出现。
//append函数扩展，其接收一个切片类型和一个可变长的参数。
slice1 = append(slice1,"Ruby","Erlang")
//append还可以将两个元素类型相同的切片值连起来。
slice1 = append(slice1,slice2...)

var array2 [10]int = [10]int{0,1,2,3,4,5,6,7,8,9}
slice5 := array2[2:6]//这样指定使得可以肯定元素的下界索引为2,其上界索引不能指定，为array2的容量上界
slice5 := array2[2:6:8]//这样就可以指定访问array2的范围为[2,8)之内的元素值。其容量为上界元素8与下界元素2的差值，即使使用slice5[:cap(slice5)]扩展也不能改变其范围。
slice5 = append(slice5,[]int{10,11,12,13,14,15}...)//这样做也不能改变其容量，使得底层数组的值改变，切断了与底层数组array2的联系。
// slice5[:3:5]//合法
// slice5[0::5]//编译错误
sliceA := []string{"Notepad","UltraEdit","Eclipse"}
sliceB := []string{"Vim","Emacs","LiteIDE","IDEA"}
n1 := copy(sliceA,sliceB)//copy作用把源切片值（第二个参数）中的元素赋值到目标切片值（第一个参数）中，并返航被赋值的元素值的数量，两个参数类型一致。且长度等于较短的那个切片长度。
//两个交换之后，n1仍为3,但是sliceB 为[]string{"Notepad","UltraEdit","Eclipse","IDEA"}

//字典操作：
//向集合中添加键值对
//从集合中删除键值对
//修改集合中已存在的键值对的值
//查找一个特定键所对应的值
//无序的，map[k]T  键的类型为k，元素类型为T 键的类型不能是函数、字典、切片。
//值表示法：键与元素之间用：隔开，元素之间用，隔开，用{}括起来。
map[string]bool {"vim":true,"Notepad":false}
//GO语言中函数调用只有传值，没有传引用，是否对其存放位置的内容改变，会取决于改变值的类型是值类型还是引用类型。
editorSign := map[string]bool {"vim":true,"Notepad":false}
sign1  := editorSign["Vim"]//如果没有“vim”则元素值为false，因此会存在歧义，解决方法：
sign2,ok := editorSign["Vim"]	//if ok为true，则存在。
delete(editorSign,"Vim") //删除某个键用delete，delete两个参数，第一个字典，第二个键。
//字典值不是并发安全的。可以扩展字典类型的安全型，需要使用标准库代码包sync中的结构体RWMutex



	fmt.Println("Hello world")
	fmt.Println(cap(slice1))
}
func (name string,age int,seniority int,informations ...string)bool{}
func (name string,age int,seniority int,informations ...string)(done bool){}
//函数的零值是nil，函数分为命名函数和匿名函数，匿名函数只是没有函数名称而已。
func Module(x,y int)(result int){
	result = x % y
	func (x,y int)(result int){
		result = x % y
		return 
	}
	return 
}//返回无参数时，result可以被传出去。

//利用函数类型定义一个变量
var recorder func(name string,age int,seniority int)(done bool)

recorder = func(name string,age int,seniority int)(done bool){
	return
}

done := recorder("Harry",32,10)

func (name string,age int,seniority int)(done bool){
	return
}("Harry",32,10)	//匿名函数使用

type Encipher func(plaintext string)[]byte	//类型替换


func GenEncryptionFunc(encrypt Encipher) func(string)(ciphertext string){
		return func(plaintext string) string{
			return fmt.Sprintf("%x",encrypt (plaintext))
		}
}//闭包  只有当函数类型是一等类型且其值可以作为其他函数的参数或结果的时候，才能编写出闭包的代码。
//方法是附属于某个自定义的数据类型的函数。
type MyIntSlice []int
func (self MyIntSlice) Max() (result int){
	return 
}//方法名为Max 接收者声明为self MyIntSlice,MyIntSlice为该方法的标识符所属的类型，
//方法规则：接收者声明中的类型必须是某个自定义的数据类型，或者是一个与某自定义数据类型对应的指针类型。
func (self *MyIntSlice) Min() (result int)
//接收者声明中的类型必须由非限定标识符代表。也就是说，方法的所属的数据类型的声明必须与该方法声明在同一个代码内
//接收者标识符不能为空标识符"_"，并且必须在所在的方法签名中唯一。
//func (self *MyIntSlice) Min() (result int)的类型是func Min() (self *MyIntSlice,result int)
//接口由方法集合代表。接口类型声明由若干个方法的声明组成。方法的声明由方法名称和方法签名构成。一个接口类型不允许出现重复的方法。
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
interface{}//空接口

type SortableStrings [3]string
func (self SortableStrings) Len() int{
	return len(self)
}
func (self SortableStrings) Less(i,j int) bool{
	return self[i] < self[j]
}
func (self SortableStrings) Swap(i,j int) {
	self[i],self[j] = self[j],self[i]
}
//有了方法实现，就可以使用sort.Interface接口了。
_,ok := interface{}(SortableStrings{}).(sort.Interface)	//需要导入包sort，验证接口类型实现的直接方法。
//一个接口类型可以被任意数量的数据类型实现，相反，一个数据类型也可以同时实现多个接口类型。
func (self SortableStrings) Sort(){
	sort.Sort(self)
}

_,ok2 := interface{}(SortableStrings{}).(Sortable)
//如果把SortableStrings改为一个指针类型，则测试改为：_,ok2 := interface{}(&SortableStrings{}).(Sortable)

ss := SortableStrings{"2","3","1"}
ss.Sort()
fmt.Printf("Sortable Strings:%v\n",ss)

//结构体
//命名结构体和匿名结构体
type Sequence struct{
	len int
	cap int
	Sortable		//匿名字段
	sortableArray sort.Interface
}

type Anonymities struct {
	T1
	*T2
	P.T3
	*P.T4
}
//4个匿名字段， T1和P.T3非指针的数据类型。隐含的名称为T1和T3，*T2和*P.T4为指针类型，隐含名称T2和T4。
//匿名字段的隐含名称也不能与所属的结构体类型中的其他字段名称重复

type Sequence struct{
	Sortable		//匿名字段
	sorted bool
}

seq Sequence
seq.Sort()

//如果Sequence中也有一个Sort(),则Sortable中的Sort()会被隐藏  //Sort()的名称签名相同

type Sequence struct{
	Sortable		//匿名字段
	sorted bool
	Sort()
}

func (self *Sequence)Sort(){
	self.Sortable.Sort()
	self.sorted = true
}

//Sort()的名称相同签名不相同，Sortable中的Sort()也会被隐藏,这个时候，必须使用该类型的Sort()方法的签名来编写调用表达式

func (self *Sequence) Sort(quicksort bool){

}
//这个时候必须这样调用
seq.Sort(true) //seq.Sort(false)
//也可以这样调用
seq.Sortable.Sort()


type List struct{
	Sequence
}

//匿名结构体
struct {
	Sortable
	sorted bool
}
//缺少type和名称

//直接将匿名结构体类型作为变量的类型
var anonym struct{
	a int
	b string
}
//以匿名结构体类型为变量初始化
anonym := struct{
	a int
	b string
}(0,"string")

//结构体类型声明中的字段声明后面添加字符串字面量
type Person struct{
	Name string 	`json:"name"`
	Age uint8 			`json:"age"`
	Address string `json:"addr"`
}

//赋值
Sequence{Sortable:SortableStrings{"3","2","1"},sorted:false}


//限制

//如果要省略其中某个或某些键值对的键，那么其他的键值对的键必须省略。
//Sequence{SortableStrings{"3","2","1"},sorted:false}	//编译错误
//多个字段值之间的顺序应该与结构体类型声明中的字段声明的顺序一致。并且不能省略掉对任何一字段的赋值。
// Sequence{sorted:false,Sortable:SortableStrings{"3","2","1"}}
// 和
// Sequence{Sortable:SortableStrings{"3","2","1"}}	//合法  因为指定变量

// Sequence{false,SortableStrings{"3","2","1"}}
// 和
// Sequence{SortableStrings{"3","2","1"}}	//不合法  变量顺序不对

//结构体中不指定字段 的值  意味着都赋予零值 
Sequence{}
//当字段名称首字母为小写字母时，只能在结构体类型声明所属的代码包中访问。
//结构体类型是值类型。和数组类型一样

type SortableStrings []string

[]string(SortableStrings{"4","5","6"})
 SortableStrings([]string{"4","5","6"})		//一样


 type GenericSeq interface{
	 Sortable
	 Append(e interface{}) bool
	 Set(index int,e interface{}) bool
	 Delete(index int) (interface{},bool)
	 ElemValue(index int) interface{}
	 ElemType() reflect.Type
	 Value() interface{}
 }

 type Sequence struct{
	 GenericSeq
	 sorted bool
	 elemType reflect.Type	//标准库代码包reflect中的Type类型
 }

 func (self *Sequence) Sort(){
	 self.GenericSeq.Sort()
	 self.sorted = true
 }

 func (self *Sequence) Append(e interface{}) bool{
	 result := self.GenericSeq.Append(e)
	 //部分代码省略
	 self.sorted = false
	 //部分代码省略
	 return result
 }

 func (self *Sequence) Set(index int, e interface{}) bool{
	result := self.GenericSeq.Set(index,e)
	//部分代码省略
	self.sorted = false
	//部分代码省略
	return result
 }

 func (self *Sequence) ElemType() reflect.Type{
	//部分代码省略
	self.elemType = self.GenericSeq.ElemType()
	return self.elemType
 }


 //指针
 //uintptr 用于存储内存地址的类型。 属于数值类型。
 //指针类型属于引用类型，零值为nil
 //包unsafe中有一个Pointer类型，可以将任意类型转化为对应的指针类型。
 f32 float32
 pointer := unsafe.Pointer(&f32)
 vptr := (*int) (pointer)

 //一个unsafe.Pointer类型可以转换为一个uintptr类型
 uptr := uintptr(pointer)
 //一个uintptr类型也可以转换为一个uintptr类型
 pointer2 := unsafe.Pointer(uptr)

 pp := &Person{"Robert",32,"Beijing,China"}
 //取内存地址
 var puptr = uintptr(unsafe.Pointer(pp))

 var npp uintptr = puptr + unsafe.Offsetof(pp.Name)	//Offsetof返回作为参数的某字段在其所属的结构体类型之中的存储偏移量。

 var name *string = (*string)(unsafe.Pointer(npp))	//获取Name字段值

 uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))


 //new函数  为值分配内存。不会初始化分配到的内存，只会清零。
//new(T) 为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回。它的类型为*T
//bytes.Buffer是一个尺寸可变的字节缓冲区。 调用new(bytes.Buffer)则是指向一个空缓冲区的指针值。
//sync中的Mutex也是new后即可用的数据类型，它的零值就是一个处于未锁定状态的互斥量。
//字面量所代表的是该类型的值，而不是指向该类型值的指针值。因此，我们在将它们与调用new函数的调用表达式做等价替换的时候，还需要在字面量的前面加取地址符&以表示指向该类型值的指针值。

//make函数 只能被用于创建切片类型、字典类型和通道类型的值。并返回一个已被初始化（非零值）的对应类型的值
//make除了接收一个表示目标类型的字面量之外，还可以接收一个或两个的额外参数。

make([]int,10,100)	//长度10,容量100

s := make([]int,10)//长度和容量都为10

make(map[string]int,100)

make(chan int,10)

m := make(map[string]int,100)
mp := &m


//new主要用于创建值类型的值。
//make仅用于创建引用类型。

struct{
	a int
	b string
}(0,"string")
//所代表的值类型是
struct{
	a int
	b string 
}
//这个类型就是一个未命名的类型，它的潜在类型和结构体类型
type Anonym struct{
	a int
	b string
}
//相同

//引用类型：切片、字典、函数、接口、指针、通道。预定义nil



ints[1],ints[2] = (ints[1] + 1),(ints[2]+1)

ints[1],_ = (ints[1] + 1),(ints[2]+1)

//1、调用返回多个结果的表达式： id,done := Save(Person{})
//2、应用于字典的索引表达式：v,ok := map["k1"]
//3、类型断言表达式：v,ok := x.(string)
//4、由接收符和通道组成的表达式：v,ok := <-ch
//平行赋值永远从左到右，即使靠右的赋值引发恐慌，左边赋值依然生效。
x := []{1,2,3}
x[0],x[0] = 1,2
x[2],x[3] = 4,5	//x[3]越界，所以赋值失败，x[2]依然有效。

//常量
//   常量只能由字面常量或常量表达式赋值。常量表达式是能够且总会在编译期间被求值的。
//布尔常量、rune常量（字符常量）、整数常量、浮点数常量、复数常量、字符串常量。
//常量可以有类型也可以无类型。
//整数字面量至少要用256个比特位来表示
//浮点数字面量的小数部分至少要用256个比特位来表示，而指数部分至少要用32个比特位来表示。
//若不能精确的表示一个整数常量，则要给出一个错误。
//若由于溢出而不能表示一个浮点数或复数常量，则要给出一个错误。
//若由于精度限制而不能表示一个浮点数或复数常量，则这个值会被四舍五入为一个相近的常量。

const untypedConstant  = 10.0		//无类型常量。

const typedConstant  int64 = 1024

//平行赋值
const tc1,tc2,tc3 int64 = 1024,-10,88

const utc1,utc2,utc3 = 6.3,false,"C"

//拆分多行
const (
	utc1 = 6.3
	utc2 = false
	utc3 = "C"
)


const (
	utc1,utc2 = 6.3,false
	utc3 = "C"
	utc4
	utc5
)
//utc4和utc5与utc3的类型一致。
//如果有一个未显式赋值的常量，那么与它同一行的常量赋值也必须都省略。
//未包含显示赋值的那一行常量声明中的常量标识符的数量必须与它上面的、最近的且包含显式赋值的那一行常量声明中的常量标识符的数量相等。

const (
	utc1,utc2，utc3  = 6.3,false，"C"
	utc4
	utc5
)//非法


const (
	utc3 = "C"
	utc1,utc2 = 6.3,false
	utc4，utc5
)//合法

 //不但可以隐含的对多个常量赋予同一个值，而且还可以更加方便的对多个变量分别赋予一系列连续的值。
 const(
	 Sunday = iota
	 Monday
	 Tuesday
	 Wednesday
	 Thursday
	 Friday
	 Saturday
 )

 //iota代表连续的、无类型的整数常量。第一次出现在以const开始的常量声明语句中的时候总会表示整数常量0
 //在iota中第二个包含它的常量声明中整数常量为1

 const x =iota
 const y = iota	//x和y都为0

 const (
	 u  = 1 << iota
	 v
	 w
 )//结果为1,2,4原因：3个常量值分别表示1 << 0 ,1 << 1, 1 << 2

 const(
	 e,f = iota,1<<iota
	 g,h	//1,2
	 i,j		//2,4
 )//iota代表的整数常量是否递增取决于是否又有一个常量声明包含了它，而不是它是否又在常量声明中出现一次。

 const(
	e,f = iota,1<<iota
	_,_
	g,h	//2,4
	i,j		//3,8
)

//平行赋值
var v1,v2 = 0,-1

//多行赋值
var (
	v1 = 0
	v2 = -1
)
//不初始化的变量声明，类型不能省略。


//短变量声明仅能够在函数体内部声明变量的时候使用。
//在短变量声明中:=的左边的标识符至少要有一个代表在当前上下文环境中的新变量。


//类型恒等。
//如果名称相同且源于相同的类型的声明，则恒等。
//别名类型和它的源类型是两个完全不同的类型，一个命名类型和一个匿名类型总是不恒等的。
//两个匿名类型的字面量相同，则就恒等。

//数组，类型和长度相同则恒等。
[4]string
[4]string	//恒等。

[4]string
[3]string //不恒等。

//切片类型，数据类型恒等。
[]string
[]string	//恒等

//结构体，字段声明的数量相同，且对应位置上的字段具有相同的字段名称和恒等的数据类型。
var a1 struct{
	f1 sort.Interface
	f2 int64
}

var a2 struct{
	f1 sort.Interface
	f2 int64
}	//恒等。


var a1 struct{
	sort.Interface
	f2 int64
}

var a2 struct{
	f1 sort.Interface
	f2 int64
}//不恒等。	有匿名字段，a2不是匿名字段。


var a1 struct{
	f1 sort.Interface	`json:"list"`
	f2 int64
}

//指针类型，基本类型恒等。

//函数类型，相同的数量的参数声明和结果声明，并且对应位置上的参数或结果类型都是恒等的。
func(name string,dept string,isManager bool) (id int,done bool)

func (appName string,targetOs string,authRequired bool) (id int,done bool)	//恒等
//关心参数的数量、顺序、类型。
//如果一个可变参数，另一个也必须可变参数

func(string ,..int) bool

func (string,[]int) bool	//不恒等

//接口类型。拥有相同的方法集合。则恒等。
//两个接口类型的方法数量必须相同，每一个方法声明能够在另一个接口找到与它完全相等的方法声明。//方法顺序无所谓，名称相同是方法恒等的一个条件。

type Ia interface{
	Name() string
	Age() int
}

type Ib interface{
	Age() int
	Name() string
}		//恒等。

//字典类型。键类型和元素类型恒等。
//通道类型。具有恒等的元素类型，方向相同。

//可比性：判断两个具有可比性值是否相等

//布尔类型：可比性。
//整数类型：可比性，有序性。
//浮点类型：可比性，有序性。
//复数类型：可比性
//字符串类型：可比性，有序性。
//指针类型：可比性
numArray := [3]int{1,23,456}
p1 := &numArray
p2 := &numArray
fmt.Printf("%v\n",p1 == p2)		//true

//通道类型：可比性	，元素类型和缓冲区大小一致。nil也相等
//接口类型：可比性，相同的动态类型和动态值。 //空值也相等
type Ic interface{
	Code() string
}

type Sc interface{
	code string
}

func(self Sc) Code() string{
	return self.code
}	//结构体Sc是接口类型Ic的一个实现类型。
var ic1 Ic = Sc{code:"A"}
var ic2 Ic = Sc{code:"A"}	
fmt.Printf("%v\n",ic1 == ic2)	//true

//非接口类型X 的值x可以与接口类型T的值t判断相等。当且仅当接口类型T具有可比性，且类型X 是接口类型T 的实现类型时。

 var sc1 Sc = Sc{code: "A"}
 fmt.Printf("%v\n",sc1 == ic1)//true
 fmt.Printf("%v\n",sc1 == ic2)//true
//因为Sc类型是Ic类型的实现类型，所以可比，又由于ic1和ic2的动态类型和动态值都分别与sc1类型和值相等。

//结构体所有字段都具有可比性。结构体的值也具有可比性。
type Sd struct{
	ints []int
}

sd1 := Sd{ints:[]int{0,1}}
sd2 := Sd{ints:[]int{0,1}}

fmt.Printf("%v\n",sd1 == sd2)	//编译错误，切片不具有可比性。

//数组类型值可比性，对应位置上的值相等。
slice1 := [3][]int{[]int{0,1}}		//[]int代表[]int的数组类型值。
slice2 := [3][]int{[]int{0,1}}
fmt.Printf("%v\n",slice1 == slice2)	//编译错误，此行出错，类型和元素值不具有可比性。

//接口可比时，动态类型不具有可比性，会引发恐慌。//例如引用或字典类型时
type Se []int

func(self Se)Code()string{
	return ""
}
//实现Ic类型的接口可以被赋值给Ic类型的变量
var ic3 Ic = Se{1,2}
var ic4 Ic = Se{1,2}
//判断ic4和ic3会引发恐慌

//又：同一个接口类型，实现类型为包含切片类型字段的结构体类型。//这样就不具有可比性。
func (self Sd) Code() string{
	return ""
}
//实现Ic类型的接口可以被赋值给Ic类型的变量
var ic5 Ic = Sd{ints:[]int{0,1}}
var ic6 Ic = Sd{ints:[]int{0,1}}	
//判断ic5和ic6会引发恐慌

//切片类型、字典类型、函数类型不具有可比性。
//可以与nil值判等。

//类型转换
T(x)//T为转换后的类型。x为表达式

*string(v)	//将v转换为string类型，再获取指向它的指针类型的值。	//等同于*(string(v))

(*string)(v)	//转换为指针类型*string的值。

<-chan int(v)	//先将v转换为chan int类型，再从此通道类型中接收一个int类型的值。//等同于<-(chan int(v))

(<-chan int)(v)	//把v转换称通道类型<-chan int类型的值。

fun()(v)	//任何无参声明的一个匿名函数

(func())(v)	//将v转换为func()类型。

fun() int(v)	//等同于(func()int)(v)

//x可以被类型T 的值代表
//x是一个浮点数常量，T是一个浮点数类型。并且x在被舍入后可以被类型T 的值代表。
//如果x是一个整数数常量，并且T是一个string类型，那么将会遵循一套规则来决定类型转换的结果。

//值x可以被赋给类型T的变量。
type Computer interface{
	CpuType() string
}

type Laptop struct{
	cpuType string
}

func(self Laptop) CpuType() string{
	return self.cpuType
}
Computer(Laptop{cpuType:"Intel Core i5"})	//合法。//因为Laptop是Computer的一个实现类型。

//值x的类型和值T的潜在类型相等。
type MyString string
MyString("Mine")	//合法。

//值x的类型和类型T都是未命名的指针类型，并且它们的基本类型的潜在类型都是相等的。
var str1  string
(*MyString)(&str1)

//值x的类型和类型T都是整数类型或都是浮点数类型。
var com64 complex64
complex128(com64)	//合法

//值x是一个整数类型值或一个元素类型为byte或rune的切片类型值，且T是一个string类型。
string([]byte{'a'})	//合法		且结果为string类型"a"

//值x是一个string类型值，且T是一个元素类型byte或rune的切片类型。
[]rune("golang")	//合法的。


//数值类型之间的转换
var number int = 1024
//或
int(1024)

//string类型相关
//byte类型的切片类型向字符串类型转换。
string([]byte{'g','o','l','a','n','g'})
//把一个数据类型rune的切片类型值向字符串类型转换。//产生一个字符串类型的值。
string([]rune{0x4e2D,0x56fd})
//把一个字符串类型向[]byte类型转换，其结果会把该字符串类型值按字节拆分后的结果。
[]byte("hello")	//结果：[]byte{104,101,108,108,111}
//把一个字符串类型的[]rune类型转换时，其结果将会是把该字符串类型值按字节拆分后的结果。
[]rune("中国")	//结果：[]rune{20013,22269}

//byte和rune都属于整数值的一种。


//内建函数
//close函数	//只接收通道类型的参数
ch := make(chan int,1)
close(ch)

//len函数和cap函数

/*
//len使用
参数类型				结果
string 							string类型字节长度
[n]T或*[n]T				数组类型值长度	n
[]T								切片类型长度
map[k]T					子典类型长度，键的数量
chan T						通道类型当前包含的元素数量

//cap使用
参数类型				结果
[n]T或*[n]T				数组类型值长度	n
[]T								切片类型容量
chan T						通道类型容量
*/

//切片类型：0 <= len(s) <= cap(s)
//len()的值在编译期间被计算出来。

//new和make

//append和copy函数
//用于辅助切片类型之上的操作。

//delete函数
//删除一个字典类型值中的某个键值对。
//两个参数：第一个目标的字段类型值，第二个要删除的那个键值对的键。
delete(m,k)
//k与m的键的类型之间必须满足赋值规则
//m为nil或k所代表的键值对并不存在于m中的时候，delete(m,k)不会做任何操作。没有删除的操作时，不做任何操作。

//complex、real和imag函数
//专门用于复数操作的。
//complex用于根据浮点数类型的实部和虚部来构造复数类型值。
var cplx128 complex128 = complex(2,-2)
//real和imag分别用于从一个复数中抽取浮点类型的实部和虚部。
var im64 = imag(cplx128)
var r64 = real(cplx128)

z == complex(real(z),imag(z))	//z是一个复数类型。

//panic和recover函数
//用于报告和处理运行时的恐慌。
//panic只接收一个参数。任意类型值。
//recover不接收任何参数。返回一个interface{}类型值。


//print和println函数
//print从左到右打印出传递给它的参数值。
//println在print基础上每个参数之间加一个空格，最后加一个换行符。


//流程控制
//if语句后必须跟{}.

if 100  < number{
	number++
}else{
	number--
}

if diff:= 10 - number; 100 < diff{
	number++
}else {
	number--
}

if diff:= 10 - number; 100 < diff{
	number++
}else if 200 < diff {
	number--
}else{
	number -= 2
}

//os包中的Open
func Open(name string) (file *File,err error)

f,err := os.Open(name)
if err != nil{
	return err
}	//必须检查err是否为nil

func update(id int,deptement string) bool{
	if id <= 0{
		return false
	}

	return true
}

//改进：
func update(id int,deptement string) bool{
	if id <= 0{
		return errors.New("The id is INVALID!")
	}

	return true
}

//switch
//表达式switch，类型switch

//表达式switch
switch content{
default:
	fmt.Println("Unkown language")
case "Python":
	fmt.Println("An interpreted Language")
case "Go":
	fmt.Println("A compiled language")
}

switch content := getContent();content{
default:
	fmt.Println("Unkown language")
case "Python":
	fmt.Println("An interpreted Language")
case "Go":
	fmt.Println("A compiled language")
}


switch content := getContent();content{
default:
	fmt.Println("Unkown language")
case  "Ruby","Python":
	fmt.Println("An interpreted Language")
case  "C","Java","Go":
	fmt.Println("A compiled language")
}
//fallthrough语句会将流程控制权转移到下一条case语句上。且在此case代码块中的最后一条语句。同时不能出现在最后一个case的语句列表中。
switch content := getContent();content{
default:
	fmt.Println("Unkown language")
case  "Ruby":
	fallthrough
case "Python":
	fmt.Println("An interpreted Language")
case  "C","Java","Go":
	fmt.Println("A compiled language")
}

switch content := getContent();content{
default:
	fmt.Println("Unkown language")
case  "Ruby":
	break
case "Python":
	fmt.Println("An interpreted Language")
case  "C","Java","Go":
	fmt.Println("A compiled language")
}

//break和标记之间要用空格隔开。


//类型switch
//利用type判断类型
switch v.(type){
case string:
	fmt.Printf("The string is '%s'.\n",v.(string))
case int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64:
	fmt.Printf("The integer is %d.\n",v)
default:
	fmt.Printf("Unsupported value.(tyep = %T\n)",v)
}
//v的类型是某个接口类型。case表达式中的类型字面量必须是v的类型的实现类型。
//通用方案：把v类型设置为interface{}类型。
//case表达式类型字面量可以为nil。
//fallthrough不允许出现在类型switch语句中任何case语句中。
switch i := v.(type){
case string:
	fmt.Printf("The string is '%s'.\n",i)	//相当于 i:=v.(string)
case int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64:
	fmt.Printf("The integer is %d.\n",i)
default:
	fmt.Printf("Unsupported value.(tyep = %T\n)",i)
}

//switch没有表达式的情况下，视为布尔类型true,其case语句也是bool类型
switch{
case number < 100:
	number++
case number < 200:
	number--
default:
	number -= 2
}

switch number := 123; {
case number < 100:
	number++
case number < 200:
	number--
default:
	number -= 2
}

//for

for number < 200{
	number += 2
}

for{
	number++
}

for i := 0;i < 100;i++{
	number++
}

var j uint = 1
for ;j %5 != 0;j *= 3{
	number++
}

for k := 1;k % 5 != 0;{
	k *= 3
	number++
}

//range子句
ints := []int{1,2,3,4,5}

for i,d := range ints{
	fmt.Printf("%d:%d\n",i,d)
}

var i,d int
for i,d = range ints{
	fmt.Printf("%d:%d\n",i,d)
}

//range表达式类型：数组、指向数组的指针值、切片、字符串、字典、通道类型。

ints := []int{1,2,3,4,5}
length := len(ints)
indexesMirror := make([]int,length)
elementsMirror := make([]int,length)

var i int
for indexesMirror[length - i - 1],elementsMirror[length - i - 1] = range ints{
	i++
}

m := map[uint]string{1:"A",6:"C",7:"B"}
var maxKey uint
for k := range m{
	if k > maxKey{
		maxKey = k
	}
}

var values []string
for _,v := range m{
	values = append(values,v)
}


var nameCount map[string]int

targetsCount := make(map[string]int)
for k,v := range nameCount{
	matched := true
	for _,r := range k{
		if r < '\u4e00' || r > '\u9fbf'{
			matched = false
			break
		}
	}
	if matched {
		targetsCount[k] = v
	}
}


for k,v := range nameCount{
	matched := true
	for _,r := range k{
		if r < '\u4e00' || r > '\u9fbf'{
			matched = false
			break
		}
	}
	if !matched {
		break
	}else{
		targetsCount[k] = v
	}
}

L:
	for k,v := range nameCount{
		if v > 100{
			fmt.Printf("The matched name:%v\n",k)
			break L				//终止L
		}
	}


	var nameCount map[string]int

	targetsCount := make(map[string]int)
	L:
	for k,v := range nameCount{
		matched := true
		for _,r := range k{
			if r < '\u4e00' || r > '\u9fbf'{
				matched = false
				break L
			}
		}
		targetsCount[k] = v
	}



for k,v := range nameCount{
	matched := true
	for _,r := range k{
		if r < '\u4e00' || r > '\u9fbf'{
			matched = false
			break
		}
	}
	if !matched {
		continue
	}else{
		targetsCount[k] = v
	}
}

var nameCount map[string]int

	targetsCount := make(map[string]int)
L:
	for k,v := range nameCount{
		for _,r := range k{
			if r < '\u4e00' || r > '\u9fbf'{
				continue L
			}
		}
		targetsCount[k] = v
	}

	for i,j := 0,len(numbers) - 1;i < j;i,j = i +1,j-1{
		numbers[i],numbers[j] = numbers[j],numbers[i]
	}

	//goto
	//1、不允许因使用goto而使当前作用也中的变量进入该作用域.
	goto L
	v := "B"
L:
	fmt.Printf("V:%v\n",v)

	//修改：
	v := "B"
	goto L
L:
	fmt.Printf("V:%v\n",v)			//因为goto跳过了v的声明

	//2、goto直属语句为A ，标记语句为B，如果B 不是A 的外层语句，就不合法。

	if n%3 != 0{
		goto L1
	}

	switch{
	case n%7 == 0:
		fmt.Printf("%v is a common multiple of 7 and 3.\n",n)
	default:
	L1:
		fmt.Printf("%v isn't a multiple of 3.\n",n)
	}

	//修改为
	if n%3 != 0{
		goto L1
	}

	switch{
	case n%7 == 0:
		fmt.Printf("%v is a common multiple of 7 and 3.\n",n)
	default:
	}

	L1:
	fmt.Printf("%v isn't a multiple of 3.\n",n)	//标记语句在switch中，不是if的外层语句。



func findEvildoer(name string)string{
	var evildoer string
	for _,r := range name{
		switch{
		case r >= '\u0041' && r <= '\u005a':	//a-z
		case r >= '\u0061' && r <= '\u007a':	//A-Z
		case r >= '\u4e00' && r <= '\u9fbf':	//中文字符
		default:
			evildoer = string(r)
			goto L2
		}
	}
	goto L3
L2:
	fmt.Printf("The first evildoer of name '%s' is '%s'!\n",name,evildoer)
L3:
	return evildoer
}


func findEvildoer(name string)string{
	var evildoer string
L2:
	for _,r := range name{
		switch{
		case r >= '\u0041' && r <= '\u005a':	//a-z
		case r >= '\u0061' && r <= '\u007a':	//A-Z
		case r >= '\u4e00' && r <= '\u9fbf':	//中文字符
		default:
			evildoer = string(r)
			break L2
		}
	}
	if evildoer != ""{
		fmt.Printf("The first evildoer of name '%s' is '%s'!\n",name,evildoer)
	}
	return evildoer
}


func checkValidity(name string) error{
	var errDetail string
	for i,r := range name{
		switch{
		case r >= '\u0041' && r <= '\u005a':	//a-z
		case r >= '\u0061' && r <= '\u007a':	//A-Z
		case r >= '\u0030' && r <= '\u0039':	//0-9
		case r == '_' || r == '-' || r == '.':	//其他允许的字符
		default:
			errDetail = "The name contains some illegal characters"
			goto L3
		}
		if i== 0{
			switch r {
			case '_':
				errDetail = "The name can not begin with a '_'."
				goto L3
			case '-':
				errDetail = "The name can not begin with a '-'."
				goto L3
			case '.':
				errDetail = "The name can not begin with a '.'."
				goto L3
			}
		}
	}
	return nil
L3:
	return error.New("Validity check failure:" +errDetail)
}


//defer  
//只能出现在函数或方法内部。
//当外围函数的函数体中的相应的语句全部被正常执行完毕的时候，只有在该函数中的所有defer语句被执行完毕之后该函数才会真正的结束执行
//当外围函数体中的return被指向的时候，只有在该函数中的所有defer都被执行完毕之后该函数才会真正的返回。
//当外围函数发生恐慌的时候，只有在该函数的所有defer都指向完毕之后该运行恐慌才会真正的被扩散到函数的调用处。

func isPositiveEvenNumber(number int) (result bool){
	defer fmt.Println("done.")
	if number < 0 {
		panic(errors.New("The number is a negative number!"))
	}
	if number % 2 == 0{
		return true
	}
	return
}

//defer:收尾任务总会被执行，我们可以把它们放到外围函数的函数体中的任何地方，而不是放在函数体的最后。
//匿名函数
defer func(){
	fmt.Println("The finishing touches.")
}()

func printNumbers(){
	for i := 0;i < 5;i++{
		defer fmt.Printf("%d",i)
	}
}
// 4 3 2 1 0

func appendNumbers(ints []int) (result []int){
	result = append(ints,1)
	defer func(){
		result = append(result,2)
	}()
	result = append(result,3)
	defer func(){
		result = append(result,4)
	}()
	result = append(result,5)
	defer func(){
		result = append(result,6)
	}()
	return result
}
//如果传入为[]int{0},则结果为[]int{0,1,3,5,6,4,2}

func printNumbers(){
	for i := 0;i < 5;i++{
		defer fun(){
			fmt.Println("%d",i)
		}()
	}
}
//5 5 5 5 5		//因为当延迟函数被逆序的逐个放入，此时i为5,因此输出5 5 5 5 5

//修改为
func printNumbers(){
	for i := 0;i < 5;i++{
		defer fun(i int){
			fmt.Println("%d",i)
		}(i)	
	}
}
//4 3 2 1 0

//如果延迟函数是一个匿名函数，并且外围函数的声明中存在命名的结果声明，那么在延迟函数代码是可以对命名结果进行访问和修改的。
func modify(n int) (number int){
	defer func(){
		number+= n
	}()
	number++
	return
}
//结果总为3
//先指向number++,再执行number += n
//虽然延迟函数声明中可以包含结果声明，但是返回的结果值会在它被指向完毕时被丢弃。
func modify(n int) (number int){
	defer func(plus int) (result int){
		result = n + plus
		number+= result
		return
	}(3)
	number++
	return
}//6


//error内建接口
type error interface{
	Error() string
}

func New(text string) error{
	return &errorString{text}
}

type errorString struct{
	s string
}

func (e *errorString) Error() string{
	return e.s
}

var err error = errors.New("A normal error");
fmt.Println(err);

err2 := fmt.Errorf("%s\n","A normal error");

//os.PathError是error接口类型的实现类型。
type PathError struct{
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string{
	return e.Op + " " + e.Path + " :" + e.Err.Error()
}

file,err3 := os.Open("/etc/profile")
if err3 != nil{
	if pe,ok := err3.(*os.PathError); ok{
		fmt.Printf("Path Error : %s (op = %s,path = %s)\n",pe.Err,pe.Op,pe.Path)
	}else {
		fmt.Printf("Uknown Error : %s\n",err3)
	}
}

r := bufio.NewReader(file)		//创建一个可读取文件内容的读取器
var buf bytes.Buffer		//缓存从文件读取出来的内容
for{
	byteArray,_,err4 := r.ReadLine()			//
	if err4 != nil{
		if err4 == io.EOF{		//io.EOF的声明为：var EOF = error.New("EOF")
			break
		}else{
			fmt.Printf("Read Error:%s\n",err4)
			break
		}
	}else {
		buf.Write(byteArray)
	}
}




type Error interface{
	error
	Timeout() bool
	Temporary() bool
}

if nerError,ok := err.(net.Error); ok && netErr.Temporary(){
	//
}


//panic

func outerFunc(){
	innerFunc()
}

func innerFunc(){
	panic(errors.New("A intended fatal error!"))
}

myIndex := 4
ia := [3]int{1,2,3}
_ = ia[myIndex]		// 数组越界

type Error interface{
	error
	RuntimeError()
}


//recover//拦截恐慌

defer func(){
	if r := recover(); r != nil{
		fmt.Printf("Recovered panic:%s\n",r)
	}
}

func (s *ss) Token(skipSpace bool,f func(rune) bool) (tok []byte,err error){
	defer func(){
		if e := recover(); e != nil{
			if se,ok := e.(scanError);ok{
				err = se.err
			}else{
				panic(e)
			}
		}
	}()
}



//完整代码：//处理恐慌：需理解
func main(){
	fethcDemo()
	fmt.Println("The main function si executed")
}

func fethcDemo(){
	defer func(){
		if v:= recover(); v != nil{
			fmt.Printf("Recovered a panic.[index = %d]\n",v)
		}
	}()
	ss := []string{"A","B","C"}
	fmt.Printf("Fetch the elements in %v one by one ...\n",ss)
	fetchElement(ss,0)
	fmt.Println("The elements fetching is done.")			//永远不会执行
}

func fetchElement(ss []string,index int) (element string){
	if index >= len(ss){
		fmt.Printf("Occur a panic ! [index=%d]\n",index)
		panic(index)
	}
	fmt.Prinf("Fetching the element...[index=%d]\n",index)
	element = ss[index]
	defer fmt.Printf("The element is \"%s\".[index=%d]\n",element,index)
	fetchElement(ss,index+1)
	return
}
//结果：
// 1、Fetch the elements in [A B C ] one by one ...
// 2、Fetching the element...[index=0]
// 3、Fetching the element...[index=1]
// 4、Fetching the element...[index=2]
// 5、Occur a panic ! [index=3]
// 6、The element is "C".[index=2]
// 7、The element is "B".[index=1]
// 8、The element is "A".[index=0]
// 9、Recovered a panic.[index = 3]
// 10、The main function si executed

//程序加文档被成为软件
//单元测试框架：java：JUnit，C++：CppUnit，Python：PyUnit。

//go test命令 以及testing包。
//测试源码文件和被它测试的源码文件在同一个代码包中。
//testing与go test命令协同使用。
//测试函数为：
func TestXxx (t *testing.T)		//t记录测试后的结果。
//参数t中的Log和Logf方法记录常规信息。
//t.Log与fmt.Println类似。
//t.Logf与fmt.Printf类型。
t.Log("Test tcp listener & sender (serverAdd=","127.0.0.1:8080",")...")
t.Logf("Test tcp listener & sender (serverAdd= %s)...","127.0.0.1:8080")
//参数t上的Error和Eroorf用于错误信息。
//当被测试的程序实体的状态不正确的时候，使用t.Error和t.Errorf，及时对当前的错误状态进行记录。
actLen := len(s)
if actLen != expLen {
	t.Errorf("Error: The length of slice should be %d but %d.\n",expLen,actLen)
}

//调用t.Error方法相当于先调用t.Log和t.fail方法进行调用。而调用t.Errorf方法则相当于先调用t.Logf和t.Fail方法调用。

//参数t上的Fatal和Fatalf方法用于记录致命的被程序实体的状态错误。
//致命错误：测试无法继续进行的错误。
if listener == nil {
	t.Fatalf("Listener startup failing!(addr = %s)!\n",serverAddr)
}//当listener没有启动时，后续工作无法执行。

//调用t.Fatal方法相当于先后调用t.Log和t.FailNow方法。调用t.Fatalf先后调用t.Logf和t.FailNow。
//标记当前测试函数测试是失败的，t.Fail方法。t.Fail方法的调用不会终止当前测试函数的执行。

//立即失败标记
//方法t.FailNow与t.Fail不同的是：它被调用时会立即终止当前测试函数的执行。会转而执行其他函数。
//t.FailNow只能在Goroutine中调用。

//失败判断
//t.Failed方法后，会得到一个bool类型的值。

//忽略测试
//调用t.SkipNow方法意味着当前测试函数已经被忽略并且立即终止该函数的执行。
//只能在运行测试函数的Goroutine中调用

//调用t.Skip相当于先后调用t.Log 和t.SkipNow。
//调用t.Skipf相当于先后调用t.Logf和t.SkipNow。
//t.Skipped结果值会告知我们当前的测试是否已被忽略。

//并行运行
//t.Parallel的调用会使当前的测试函数被标记为可并行运行。

//go test cnet/ctcp
//go test tcp_test.go tcp.go basic.go
// go test -run=Prime cnet/ctcp
//  -v:冗长模式。作用：打印的测试被记录到日志中。

//go test -timeout 时间 文件
//在规定的时间内完成测试，如果没有完成，则发生恐慌。-timeout标记的值的类型是time.Duration可接受的时间表示法。

//不设置上限，尽快结束。-short

if testing.Short() {
	multiSend(serverAddr,"SenderT",1,(2*time.Second),showLog)
} else {
	multiSend(serverAddr,"SenderT",2,(2*time.Second),showLog)
	multiSend(serverAddr,"SenderT",1,(2*time.Second),showLog)
}

//-parallel  多核并行标记

//基准测试
//基准测试函数签名
func BenchmarkXxx(b *testing.B)
//testing.B也有Log*,Error*,Fatal*,Fail*,Skip*几种方法。
//定时器方法：StartTimer,StopTimer,ResetTimer。
//go test -bench="." -v
//-benchtime  t 规定执行上限时间
//-benchmem 内存分配统计信息
//-bench regexp  regexp可以被替换为正则表达式，可以这样写：-bench.  或-bench=.
//go test -bench=Prime  cnet/ctcp
//go test -bench=Prime -benchtime 20s cnet/ctcp
// -cpu  可以是一个整数列表，  
//设置go语言最大并发处理数时，也就是调用runtime.GOMAXPROCS函数把对应整数传入。
//go test -bench=Prime -cpu=1,2,4,8,12,16,20 cnet/ctcp
//-parallel


//样本测试
//样本测试函数
func ExampleXxx() 

//样本测试命名
//1、整个代码包样本函数名称应该是Example。
//2、被测试对象是一个函数时，对于函数F，则名称为ExampleF
//3、被测试对象是一个类型时，对于类型T，则名称为ExampleT
//4、被测试对象是某个类型的一个方法时，对于类型T的方法M，则名称为ExampleT_M
//5、测试函数名称上添加后缀通过_后缀与名称隔离开。后缀为basic 则ExampleT_M_basic


//收集资源使用情况
//-cpuprofile cpu.out     cpu.out为指定文件名
//-memprofile mem.out  mem.out为指定文件名
//-memprofilerate  n   n代表分析器取样的间隔，单位字节。
//go test -cpuprofile cpu.out net		//net标准库的代码包

//go test -cpuprofile cpu.out net runtime

//go tool pprof ./net.test $GOROOT/src/pkg/net/cpu.out

//go test -memprofile mem.out -memprofilerate 10 net

//go tool pprof ./net.test $GOROOT/src/pkg/net/mem.out

//GOGC //go的垃圾回收器


//-blockprofile block.out
//-blockprofilerate n

//go test -blockprofile block.out -blockprofilerate 100 net


//测试覆盖率
//-cover		启动覆盖测试分析
//-covermode=set		set记录是否被执行过，count记录执行的次数，atomic记录执行次数，并保证并发执行正确次数。	
//-coverpkg bufio,net
//-coverprofile cover.out		

//go test -cover cnet/ctcp

//cover接受的标记
//-func=cover.out
//-html=cover.out
//-mode=count
//-o=cover.out
//-var=GoCover

//程序文档
//godoc查看本机工作文档
godoc -http=:9090 -index
//代码包注释放在doc.go文件中。



//并发编程
//多进程编程
//进程间通信方式（IPC)
//从处理机制角度来看：分为三大类：基于通讯的IPC 方法、基于信号的IPC方法、基于同步的IPC方法。
//基于通讯的：分为以数据传送为手段的IPC 方法（管道（传送字节流）、消息队列（传送结构化的消息对象的））和以共享内存为手段的IPC 方法（共享内存区，最快的IPC方法）。
//基于信号的IPC 方法：信号机制。（唯一异步IPC方法）
//基于同步的IPC方法：（信号灯）
//go支持的管道、信号、Socket。

//进程：一个程序的执行或者说用于描述程序的执行过程。
//子进程复制父进程的数据段、堆栈副本，共享代码段。
//进程id为1的为内核启动进程。
//os包中提供查看当前进程的PID和PPID。
pid := os.Getpid()
ppid := os.Getppid()

//进程的状态：运行态、可中断的睡眠态、不可中断的睡眠态、暂停态或跟踪态、僵尸态、退出态。
//虚拟内存的容量与实际可用的物理内存的大小无关，内核和cpu负责维护虚拟内存和物理内存之间的映射关系。
//内核为每个用户分配的是虚拟内存，不是物理内存。
//内核把进程的虚拟内存划分为若干页，物理内存单元的划分由cpu负责。
//一个物理内存单元称为一个页框。
//原子操作：执行中不能被中断。
//只能串行化访问或执行的某个资源或某段代码称为临界区。
//原子操作go：sync/atomic
//保证只有一个进程或线程在临界区之内的做法称为互斥。

//匿名管道
//半双工（单向）通讯方式。
//用于父子之间以及祖先的子进程之间通讯。
//简单，只能单向，通讯双方限制。
//os/exec中管道API 
cmd0 := exec.Command("echo","-n","My first command from golang")
if err := cmd0.Start(); err != nil {
	fmt.Printf("Error:The command No.0 can not be startup:%s\n",err)
	return
}
//创建一个能够获取此命令的输出管道，在上面的if之前加入：
stdout0,err := cmd0.StdoutPipe()				//stdout0类型为io.ReadCloser
if err != nil {
	fmt.Printf("Error:The command No.0 can not be startup:%s\n",err)
	return
}

output0 := make([]byte,30)
n,err := stdout0.Read(output0)			//把stdout0的值读出，存储到output0,返回一个n其实际长度，err为error类型的错误。
if err != nil {
	fmt.Printf("Error:The command No.0 can not be startup:%s\n",err)
	return
}
fmt.Printf("%s\n",output0[:n])	//把output0的数据全部输出。


var outputBuf0 bytes.Buffer
for {
	tempOutput := make([]byte,5)
	n,err := stdout0.Read(tempOutput)
	if err != nil {
		if err == io.EOF{			//io.EOF 判断是否到结尾。
			break		
		} else {
			fmt.Printf("Error:Can not read data from the pipe:%s\n",err)
			return
		}
	}
	if n > 0 {
		outputBuf0.Write(tempOutput[:n])		//Write把参数值tempOutput的值写入outputBuf0的值中。
	}
}
fmt.Printf("%s\n",outputBuf0.String())


outputBuf0 := bufio.NewReader(stdout0)
output0, _, err := outputBuf0.ReadLine()		//第二个返回值为bool类型代表是否读完。如果为false，依然可以利用for读完。
if err != nil {
	fmt.Printf("Error:Can not read data from the pipe : %s\n",err)
	return
}
fmt.Printf("%s\n",string(output0))

cmd1 := exec.Command("ps","aux")
cmd2 := exec.Command("grep","apipe")

stdout1, err := cmd1.StdoutPipe()		//cmd1输出
if err != nil {
	fmt.Printf("Error: Can not obtain the stdout pipe for command: %s\n",err)
	return
}
if err := cmd1.Start();  err != nil {
	fmt.Printf("Error: The command can not be startup: %s\n",err)
	return
}

outputBuf1 := bufio.NewReader(stdout1)
stdin2,err := cmd2.StdinPipe()		//返回值：1、与该命令连接的输入管道，2、io.WriteClose接口类型的值。
if err != nil {
	fmt.Printf("Error: can not obtain the stdin pipe for command : %s\n",err)
	return
}
outputBuf1.WriteTo(stdin2)

//得到结果
var outputBuf2 bytes.Buffer
cmd2.Stdout = &outputBuf2			//命令cmd2被启动后的所有输出内容都会写入到缓冲区
if err := cmd2.Start();err != nil {
	fmt.Printf("Error: The command can not be startup: %s\n",err)
	return
}
err =  stdin2.Close()
if err != nil {
	fmt.Printf("Error: Can not close the stdio pipe : %s\n",err)
	return
}

//需要等待cmd2运行结束之后，再查看缓冲区outputBuf2的内容
if err := cmd2.Wait(); err != nil {
	fmt.Printf("Error: Can not wait for the command : %s\n",err)
	return
}


//命名管道：任何进程都可以通过命名管道交换数据
//os.Pipe()创建管道
reader,writer ,err := os.Pipe()		//返回值：1、管道输出端的*os.File类型值，2、管道输入端的*os.File类型值，

n,err := writer.Write(input)
if err != nil {
	fmt.Printf("Error: Can not write data to the named pipe: %s\n",err)
}
fmt.Printf("Written %d byte(s).[file-based pipe]\n",n)
// 和
output := make([]byte,100)
n,err := reader.Read(output)
if err != nil {
	fmt.Printf("Error: Can not read data from the named pipe:%s\n",err)
}
fmt.Printf("Read %d byte(s).[file-based pipe]\n",n)
//并发运行。命名管道默认其中一端未就绪的时候阻塞另一端的进程。
//顺序的执行则会被永远阻塞在n,err := writer.Write(input)和n,err := reader.Read(output)
//命名管道可以多路复用，当多个输入端写的时候需要考虑原子性，操作系统提供的管道不具有原子操作。
//go提供了一个被存于内存中的、有原子操作的保证的管道：
reader,writer := io.Pipe()	//返回参数：1、该管道输出端的*PipeReader类型值，2、该管道输入端*PipeWriter类型值。
//*PipeReader类型只能使用Read方法从管道中读取数据，*PipeWriter类型只能使用Write方法向管道写入数据。

//信号
//os.Signal 的声明
type Signal interface{
	String() string
	Signal()	//无意义，只是一个标识。Singal()实现都是空方法
}

//用于把操作系统发给当前进程的指定信号通知该函数调用
func Notify(c chan<- os.Signal,sig ...os.Signal)		
//参数：1、通道类型，只能传os.Signal类型的值。只能向该通道类型值放入信号值，不能取信号值。
//2、可变长参数，第一个参数后可以再附加任意个os.Signal类型的参数值
sigRecv := make(chan os.Signal, 1)
sigs := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
signal.Notify(sigRev,sigs...)
for sig := range sigRev {
	fmt.Printf("Received a signal: %s\n",sig)
}

//恢复针对它们的系统默认操作
//os/signal 中的Stop操作
func Stop(c chan<- os.Signal)  //取消之前调用的signal.Notify函数
//取消副作用
signal.Stop(sigRecv)
close(signRecv)	//利用close()函数取下副作用。


//第一阶段代码：
sigRecv1 := make(chan os.Signal,1)
sigs1 := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
fmt.Printf("Set notification for %s... [sigRecv1]\n",sigs1)
signal.Notify(sigRecv1,sigs1...)

sigRecv2 := make(chan os.Signal,1)
sigs2 := []os.Signal{syscall.SIGQUIT}
fmt.Printf("Set notification for %s... [sigRecv2]\n",sigs2)
signal.Notify(sigRecv2,sigs2...)

//第二阶段代码：	并发执行，执行完之后退出。
var wg sync.WaitGroup
wg.Add(2)
go func(){
	for sig := range sigRecv1 {
		fmt.Printf("Received a signal from sigRecv1:%s\n",sig)
	}
	fmt.Printf("End.[sigRecv1]\n")
	wg.Done()
}()

go func() {
	for sig := range sigRecv2 {
		fmt.Printf("Received a signal from sigRecv2:%s\n",sig)
	}
	fmt.Printf("End.[sigRecv2]\n")
	wg.Done()
}

//第三阶段代码：
fmt.Println("Wait for 2 seconds...")
time.Sleep(2 * time.Second)
fmt.Printf("Stop notification...")
signal.Stop(sigRecv1)
close(sigRecv1)
fmt.Printf("done.[sigRecv1]\n")

wg.Wait()


//goc2p/src/mutiproc/signal 先构建，在运行：go build ，./
//原因：go run执行一系列的操作：依赖查找、编译、打包、链接。 生成一个临时执行文件然后在执行执行文件与go build 和./ 是一样的。
//但是，为了执行mysianal执行文件产生的一个全新的进程，它与go run那个进程无关。两个进程独立，拥有自己的id
//


//socket

func Listen(net,laddr string) (Listener,error)

//函数net.Listen是为了获取一个监听器。接收两个string参数，第一个参数以何种协议监听，必须是面向字节流的协议。
//tcp，tcp4,tcp6,udp，udp4,udp6,unix(tcp在本机中通讯),unixgram(udp在本机中通讯),unixpacket (sctp在本机中通讯)
//因此第一个参数必须是tcp,tcp4,tcp6,unix,unixpacket中一个。
//第二个参数：当前程序中的网络标识。格式host:port

listener,err := net.Listen("tcp","127.0.0.1:8085")		//返回值：1、获取监听器。2、error类型

conn,err := listener.Accept()		//返回值：1、当前tcp连接的net.Conn类型值。2、error类型

//用于向网络中的某个地址发生数据
func Dial(network,address string) (Conn,error)	//接收两个参数：1、不一定建立连接，因此还可使用的协议：udp，udp4,udp6,ip，ip4,ip6,unixgram
//第二个参数：当前程序中的网络标识。格式host:port
conn,err := net.Dial("tcp","127.0.0.1:8085")

//设定超时时间函数
func DialTimeout(network,address string,timeout time.Duration) (Conn,error)
//与Dial的唯一区别就是可以设定超时时间。
//类型time.Duration单位纳秒
//常量time.Nanosecond代表1纳秒
//常量time.Microsecond代表1微妙
//2秒：2*time.Second

conn,err := net.DialTimeout("tcp","127.0.0.1:8085",2*time.Second)

//socket编程在底层是非阻塞式的。

//net.Conn类型是一个接口类型。包含8个方法，定义了在一个连接上做的所有事情。

//Read方法用来从Socket的接收缓冲区读取数据
Read(b []byte) (n int,err error)
//参数b为[]byte类型，用来存放从连接上接收到的数据的容器。长度由应用程序决定。
b := make([]byte,10)
n,err := conn.Read(b)
content := string(b[:n])

var dataBuffer bytes.Buffer
b := make([]byte,10)
for {
	n,err := conn.Read(b)
	if err != nil {
		if err == io.EOF{
			fmt.Println("The connection is closed.")
			conn.Close()
		}else {
			fmt.Printf("Read Error :%s\n",err)
		}
		break
	}
	dataBuffer.Write(b[:n])
}

reader := bufio.NewReader(conn)

line,err := reader.ReadBytes('\n')


//Write方法用来向Socket发送缓冲区写入数据
Wrtie(b []byte) (n int, err error)

//该方法屏蔽了许多非阻塞式Socket的接口细节。

writer := bufio.NewWriter(conn)

//以Write前缀的函数的参数值数据的字节数量超出了此容量，则会试图将数据全部或一部分写入到它代理的对象中。
//因此，为解决此类问题，可以使用bufio.NewWriterSize 函数初始化一个缓冲写入。


//Close方法关闭当前连接。
//不接收任何参数，返回一个error类型值。


//LocalAddr 和RemoteAddr方法都不接收任何参数并返回一个net.Addr类型的结果。
//LocalAddr返回本地地址
//RemoteAddr返回远程地址
//net.Addr类型是一个接口类型。有两个方法：Network,String
//Network返回当前连接所使用的协议名称
conn.LocalAddr().Network()
conn.RemoteAddr().String()
conn.LocalAddr().String()

//SetDeadline,SetReadDeadline，SetWriteDeadline  三个方法值接收一个time.Time类型值，返回一个error值。

//SetDeadline方法设定当前连接上的I/O 操作的超时时间。

b := make([]byte,10)
conn.SetDeadline(time.Now().Add(2 * time.Second))
for {
	n,err := conn.Read(b)
}

//超时2秒，for循环如果迭代几次，就2秒了，则下一次就会超时。改为：
b := make([]byte,10)
for {
	conn.SetDeadline(time.Now().Add(2 * time.Second))
	n,err := conn.Read(b)
}

conn.SetDeadline(time.Time{})	//取消超时设定时间

const(
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELTMITER = '\t'
)

func printLog(format string,args ...interface{}) {
	fmt.Printf("%d : %s",logSn,fmt.Sprintf(format,args...))
	logSn++			//包级私有，日志序号
}

func handleConn(conn net.Conn){
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq,err := read(conn)
		if err != nil {
			if err == io.EOF{
				printLog("The connection is closed by another side.(Server)\n")
			}
			else {
			printLog("Read Error:%s(Server)\n",err)
			}
			break
		}
		printLog("Received request :%s(Server)\n",strReq)
	}

	for {
		i32Req,err := coverToInt32(strReq)		//转换为int32类型
		if err != nil {
			n,err := write(conn,err.Error())
			if err != nil {
				printLog("Write Error (written %d bytes):%s(Server)\n",err)
			}
			printLog("Sent response(written %d bytes):%s (Server)\n"n,err)
			continue
		}
		f64Resp := cbrt(i32Req)		//计算立方根
		respMsg := fmt.Sprintf("The cube root of %d is %f.",i32Req,f64Resp)
		n,err := write(conn,respMsg)
		if err != nil {
			printLog("Write Error :%s(Server)\n",err)
		}
		printLog("Sent response(written %d bytes):%s(Server)\n",n,respMsg)
		defer conn.Close()
	}
	
}

func read(conn net.Conn) (string,error) {
	readBytes := make([]byte,1)		//初始化为1,防止从连接值中读出多余的数据从而对后续的读取操作造成影响。
	var buffer bytes.Buffer
	for {
		_,err := conn.Read(readBytes)
		if err != nil {
			return "",err
		}
		readByte := readBytes[0]
		if readBytes == DELTMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(),nil
}

func write(conn net.Conn,content string) (int,error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELTMITER)
	return conn.Write(buffer.Bytes())
}


func serverGO(){
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK,SERVER_ADDRESS)
	if err != nil {
		printLog("Listen Error :%s\n",err)
		return
	}
	defer listener.Close()
	printLog("Got listener for the server,(local address :%s)\n",listener.Addr())

	for {
		conn,err := lisener.Accept()		//阻塞直至新连接到来
		if err != nil {
			printLog("Accept Error:%s\n",err)
		}
		printLog("Established a connection with a client application.(remote address:%s)\n",conn.RemoteAddr())
		go handleConn(conn)
	}
}

func clientGo(id int){
	conn,err := net.DailTimeout(SERVER_NETWORK,SERVER_ADDRESS,2 * time.Second)
	if err != nil {
		printLog("Dial Error :%s(Client[%d])\n",err,id)
		return
	}
	defer conn.Close()
	printLog("Connected to server.(remote address:%s,local address:%s)(Client[%])\n",conn.RemoteAddr(),conn.LocalAddr(),id)
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	 for i:= 0;i < requestNumber;i++ {
		 i32Req := rand.Int31()
		 n,err := write(conn,fmt.Sprintf("%d",i32Req))
		 if err != nil {
			 printLog("Write Error:%s(Client[%d])\n",err,id)
			 continue
		 }
		 printLog("Sent request (written %d bytes):%d (Client[%d])\n",n,i32Req,id)
	 }
	 for j := 0;j < requestNumber;j++ {
		 strResp,err := read(conn)
		 if err != nil {
			 if err == io.EOF {
				 printLog("The connection is closed by another side.(Client[%d])\n",id)
			 } else {
				 printLog("Read Error: %s (Client[%d])\n",err,id)
			 }
			 break
		 }
		 printLog("Received response:%s(Client[%d])\n"strResp,id)
	 }
}

//为了让main函数等待serverGo函数和clientGo函数结束再结束。则：
var wg sync.WaitGroup

func main(){
	wg.Add(2)		//参数2,一个服务端一个客户端
	go serverGo()		//并行
	time.Sleep(500 * time.Millisecond)
	go clientGo()
	wg.Wait()		//等待
}
//执行main函数时将永远阻塞在wg.Wait()
//再serverGo和clientGo前加
defer wg.Done()
//这样main函数即可从wg.Wait()处继续执行下去。



//多线程

//线程
//线程有自己的id叫线程ID或TID ，系统内可不唯一，但在所属的进程内唯一。
//创建线程、终止线程、连接已终止的线程、分离线程。
 //exit和return都会使其所属进程中的所有线程立即终止，还会结束该线程的运行。
 //线程的实现模型：用户级线程模型（多对一的线程实现）、内核级线程模型（一对一的线程实现）和两级线程模型（多对多的线程实现）。
 //解释：多对一是多个线程对应一个内核调度实体;一对一是一个线程对应一个内核调度实体;多对多是多个线程对应多个内核调度实体.（M:N,M和N并不一定相同）
 //线程同步：临界区、原子操作、互斥量、条件变量等。
 //共享数据一致。
 //在同一时刻，只允许一个线程处于临界区之内的约束被成为互斥。
 //每一个线程在进入临界区之前都必须先锁定某个对象。只有成功锁定对象的线程才会被允许进入到临界区之内，否则就会被阻塞。这个对象被称为互斥对象或互斥量。
 //两种状态：已锁定状态和未锁定状态。
 //互斥量每次只能被锁定一次。也就是说，处于已锁定的状态的互斥量不能被锁定。除非它已被解锁。
 //成功锁定互斥量的线程会成为该互斥量的所有者。只有互斥量的所有者才能对该互斥量进行解锁。
 //两个线程使用同一个互斥量时，只初始化一个互斥量。必须保证唯一性。
 //线程离开临界区时必须及时的解锁互斥量。


 //条件变量
 //在对应的共享数据发生状态变化时，通知其他因此而阻塞。
 //与互斥量组合使用
 //互斥量为共享数据的访问提供互斥支持，条件变量共享数据变化时向相关线程发送通知。
 //使用条件变量之前创建和初始化它。
 //初始化保证唯一性。
 //在一个条件变量之上进行3个操作：
 //1、等待通知：阻塞当前线程，直到该条件变量发来通知。
 //2、单发通知：让条件变量向至少一个正在等待它的通知的线程发送通知，以表示某个共享数据的状态已经改变。
 //3、广播通知：给等待它的所有线程发送通知。
  //函数可重入是实现线程安全最直接的方式。最简单高效的。
  
  //并发运行是指多个任务被同时发起运行，但是在同一时刻这些任务不一定都处于运行状态。
  //并行运行是指在同一时刻可以有多个任务真正的同时运行。
 //吞吐量：在一个时间单元之内程序完成并输出结果的计算任务的数量。
 //控制临界区的纯度。
 //控制临界区的粒度。
 //减少临界区中代码的执行耗时。
 //避免长时间持有互斥量
 // 优先使用原子操作而不是互斥量。

 //go并发编程。
 //两级线程模型。
 //Goroutine可以被看作是go语言特有的应用程序线程。
 //M：Machine。代表了一个内核线程。
 //P：Processor。代表了M所需的上下文环境。
 //G：Goroutine。代表了对一段需要被并发执行的Go语言代码的封装。
 //G需要M和P的支持。
//每个P包含一个可运行的G的队列。
//M和KSE(内核调度实体)是一对一的关系。
//M与P的关系也总是一对一的。P和G之间则是一对多的。
//M中有多个字段，其中字段curg会存放当前M正在运行的那个G的指针。字段p会指向与当前M相关联的那个P。
//字段mstartfn代表M的起始函数。字段nextp成为对M和P的预联。
//P本身的状态：
//Pidle:此状态表明当前P未与任何M存在关联。
//Prunning：此状态表明前P正在与某个M关联。
//Psyscall：此状态表明当前P中的被运行的那个G正在进行系统调用。
//Pgcstop：此状态表明运行时系统正在进行垃圾回收。在运行时系统进行垃圾回收的时候，会试图把全局P列表中的都置于此状态。
//Pdead:此状态表明当前P已经不会再被使用。

//Gidle:在当前G被创建但还完全未被初始化的时候会处于此状态。
//Grunnable:表示当前G是运行的，并且正在等待被运行。
//Grunning：当前G正在被运行。
//Gsyscall：当前G正在运行系统调用
//Gwaiting：当前G正在因某个原因而等待。
//Gdead:当前G已被运行完成。

//gcwaiting:表示垃圾回收任务是否正在准备或执行。0：否，1：是。
//nmidlelocked：表示已锁住且空闲的M的数量。
//nmspinning：表示正在自旋的M的数量。
//stopwait：表示在垃圾回收准备期间还未被停止的P的数量。
//sysmonwait:表示系统检测器是否正在被阻塞。0：否，1：是。

//go语句意味着一个函数或方法的并发执行。
//go语句由go关键字和表达式组成的。
go println("Go Goroutine!")

//go 匿名函数
go func() {
	println("Go Goroutine!")
}()	//这个圆括号代表了对匿名函数的调用行为。
//go之后的表达式不能被圆括号括起来。

func main() {
	go println("Go Goroutine!")
}
//上面这样并不一定会输出Go Goroutine!，因为并发运行时，会把go之后的语句放入G的队列中，至于什么时候运行，由调度器决定。
//有可能main函数执行完毕，还没有执行，是没有先后顺序的，如果要想先执行go表达式，需要处理：
func main() {
	go println("Go Goroutine!")
	time.Sleep(time.Millisecond)
}
//这样，暂时让main函数的Goroutine先停下来，进入Gwaiting状态。但是这样也不能保证go表达式一定被先执行，调度器是无法控制的。
//因此，在这里我们需要用runtime.Gosched()来代替time.Sleep()
//runtime.Gosched()作用让其他Goroutine有机会被运行。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello,%s\n",name)
	}()
	name = "Harry"
	runtime.Gosched()
}
//结果多次为Hello,Harry  因此是执行到runtime.Gosched()语句，go语句才执行。

func main() {
	names := []string{"Eric","Harry","Robert","jim","Mark"}
	for _,name := range names {
		go func () {
			fmt.Printf("Hello,%s\n",name)
		}()
	}
	runtime.Gosched()
}
//结果：
//Hello,Mark
//Hello,Mark
//Hello,Mark
//Hello,Mark
//Hello,Mark

//原因是for循环执行完之后，5个被分配的Goroutine才被执行。

//改进结果
func main() {
	names := []string{"Eric","Harry","Robert","jim","Mark"}
	for _,name := range names {
		go func () {
			fmt.Printf("Hello,%s\n",name)
		}()
		runtime.Gosched()
	}
}
//虽然这样可以得到结果，但是使得go语句变得非常复杂，因此，稍加修改。
func main() {
	names := []string{"Eric","Harry","Robert","jim","Mark"}
	for _,name := range names {
		go func () {
			time.Sleep(10 * time.Nanosecond)
			fmt.Printf("Hello,%s\n",name)
		}()
	}
	runtime.Gosched()
}
//但这样同样有问题，有可能会有一个或几个人被多次问候或没有问候。
//继续修改：
func main() {
	names := []string{"Eric","Harry","Robert","jim","Mark"}
	for _,name := range names {
		go func (who string) {
			fmt.Printf("Hello,%s\n",who)
		}(name)
	}
	runtime.Gosched()
}
//这种方案的可行性，根本原因是传入的参数是值传递，这样可以会有一个临时拷贝，如果指针传递会根据外面的变化而变化。


//main函数的Goroutine是创建的第一个Goroutine为主Goroutine，主Goroutine所做的事情并不是执行main函数，而是设定每一个
//Goroutine所能申请的栈空间的最大尺寸，32位下，最大尺寸250MB。64位下，最大尺寸1G。
//之后，主Goroutine启动系统检测器。：对调度器的工作查漏补缺。

//runtime.GOMAXPROCS函数 ：应用程序在运行期间设置运行时系统中P的最大数量。最好在Go程序之前设置。P最大数量在1-256范围内。
//runtime.Goexit函数：立即调用它的Goroutine的运行终止。其他的不受影响。
//runtime.Gosched函数：暂停调用它的Goroutine的运行。
//runtime.NumGoroutine函数：返回运行时系统中的处于特定状态的Goroutine的数量。
//runtime.LockOSThread函数:使调用它的Goroutine与当前运行它的M锁定在一起。
//runtime.UnlockOSThread函数：解锁调用它的Goroutine与当前运行它的M。
//runtime/debug.SetMaxStack函数：约束单个Goroutine所能申请的栈空间的最大尺寸。接收一个int参数为最大尺寸。
//runtime/debug.SetMaxThreads函数：对运行时系统所使用的内核线程的数量进行设置。接受一个int参数，返回一个int参数，前者欲设定的新值，后者代表之前设定的旧值。


//垃圾收集：确定垃圾并记录位置，垃圾清扫：清除垃圾并把内存归还给操作系统。
//runtime.GC函数：系统进行一次强制性垃圾收集。
//runtime/debug.SetGCPercent函数:设置垃圾收集比率。触发垃圾收集的堆内存单元增量=上一次垃圾收集完成后的堆内参单元数量*垃圾收集比率/100
//runtime/debuh.FreeOSMemory函数：手动进行一次垃圾清扫。

//happens before  原则：
//读操作r未先于写操作w发生。
//没有其他对此变量的写操作后于写操作w且先于读操作r发生。

//写操作w一定要发生在读操作r之前。
//任何其他对共享变量v的写操作都只能发生在w之前或r之后。

//channel
//通道：某一个通道类型的值。
//通道是引用类型。
//声明：
chan  T
//chan是通道关键字，T为类型。
type IntChan chan int
//int类型的通道
var intChan  chan int
//通道类型是双向的。

//单向通道：<- 用于接收操作符。

chan<- T		//只能向此类通道发送元素值。
<-chan T		//只能在此类通道中接收元素值。

<-chan int(v) //:chan int是一个整体，先将变量v的类型转换为chan int，然后再从该值中接收一个元素值。 相当于：<-(chan int(v))

//将变量v转换成一个接收通道类型：
(<-chan int) (v)

//通道为引用类型，零值为nil。
//通道类型是用来传递值的，因此它的值为即时性，是无法用字面量来准确表达的。
//相当于一个队列，已经接收的元素会被删除。具有原子性。不可分割的。

//初始化
make (chan int,10)

strChan := make(chan string,3)

//接收元素
elem := <-strChan

elem,ok := <-strChan

strChan <- "a"	//向通道内发送一个字符。
strChan <- "b"
strChan <- "c"

type Addr struct {
	city string
	district string
}

type Person struct {
	Name string
	Age uint8
	Address Addr
}

var personChan = make(chan Person,1)

p1 := Person{"Harry",32,Addr{"Beijing","Haidian"}}
fmt.Printf("p1 (1):%v\n",p1)
personChan <- p1
p1.Address.district = "Shijingshan"
fmt.Printf("p1(2):%v\n",p1)

p1_copy ：= <-personChan
fmt.Printf("p1_copy:%v\n",p1_copy)

close(personChan)	//关闭通道


func main() {
	ch := make(chan int,3)
	sign := make(chan byte,2)
	go func() {
		for i := 0;i < 5;i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
		for {
			e,ok := <-ch
			fmt.Printf("%d(%v)\n",e,ok)
			if !ok {
				break
			}	
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	<-sign
	<-sign
}