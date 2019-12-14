//使用内建函数make也可以使用map关键字定义Map
/*
	//声明变量，默认map是nil
	var map_var map[key] value_type

	//使用make函数
	map_var := make(map[key_data_type]value_data_type)

	如果不初始化就会创建一个nil map。 nil map不能用来存放键值对。

*/

package main

import "fmt"

func main(){
	var countryCapitalMap map[string]string //创建集合
	countryCapitalMap = make(map[string]string)

	//map插入key-value对，各个国家对应的首都
	countryCapitalMap ["France"] = "Paris"
	countryCapitalMap ["Italy"] = "罗马"
	countryCapitalMap ["Japan"] = "东京"
	countryCapitalMap ["India"] = "新德里"

	//使用键输出地图值
	for  country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	//查看元素在集合中是否存在
	//如果确定是真实的，则存在，否则不存在
	captial, ok := countryCapitalMap["美国"]

	if(ok){
		fmt.Println("美国的首都是：",captial)
	}else{
		fmt.Println("美国的首都不存在")
	}

}