package main

import  "fmt"
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

	