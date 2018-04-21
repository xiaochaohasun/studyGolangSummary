package main

import "fmt"

func main() {

	//申请一个指向int类型数据的指针
	//var p *int
	//*p = 10

	//打印int类型数据的内存地址及值
	var a = 10
	var p = &a
	var b = true

	fmt.Println(&a) //0xc420012058

	fmt.Println(*p)  //10
	fmt.Println(*&a) //10

	fmt.Println(*&b) //true

	//指针使用分两种情况
	//1.修改指针p指向的值，使用*p


	//2.修改指针p的地址，使用p

}
