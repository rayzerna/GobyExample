package main
/*
Go支持指针（poiter），允许你在程序中向值或记录传递引用
*/
import "fmt"

/*
我们通过2个函数：zeroval和zeroptr；进行对比，来展示指针是如何工作的。
zeroval函数有一个int形式参数（下称形参），所以传递的实际参数（下称实参）是值。
zeroval函数将从外部调用中拿到各种不同的ival副本。
*/
func zeroval(ival int) {
	ival = 0
}

/*
相比较下zeroptr函数有一个*int形参，意为数据类型为int的指针。
函数体中的*iptr会把指针解引用，也就是指向指针指向的内存地址的当前值，
并向解指引的指针分配一个新值改变原指引内存地址的值。
*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// “&i”将给出参数i的内存地址，也就是一个到i的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指标也可以打印
	fmt.Println("pointer:", &i)
}

/*
zeroval函数没有在主函数中改变i的值，但zeroptr函数改变了，
这是因为后者的变量有一个到内存的指引。
*/
//$ go run pointers.go
//initial: 1
//zeroval: 1
//zeroptr: 0
//pointer: 0x42131100
