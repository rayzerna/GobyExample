package main

/*
Go的structs是数据类型组成的集合。
它将数据组成表格形式的记录，非常有实用性。
*/
import "fmt"

// person结构体类型有name和age两个字段
type person struct {
	name string
	age  int
}

// newPerson使用形参name创建一个新person结构体
func newPerson(name string) *person {
	// 你可以安全地返回指向局部变量的指针，因为局部变量在函数作用域内还继续存在
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// 创建新结构体
	fmt.Println(person{"Bob", 20})

	// 在初始化结构体时给字段赋值
	fmt.Println(person{name: "Alice", age: 30})

	// 被忽略的字段默认为0
	fmt.Println(person{name: "Fred"})

	// 带“&”符号的配置，将产生一个指向结构体的指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 通常在创建函数中封装新结构体的创建过程
	fmt.Println(newPerson("Jon"))

	// 用点（.）访问结构体
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以对结构体指针使用点“.”————指针会自动解引用
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变的
	sp.age = 51
	fmt.Println(sp.age)
}
