package main

import "fmt"

func main() {
	//int
	//i := 1
	var i int
	i = 1
	fmt.Println(i)

	//string
	s := "hello world!"
	fmt.Println(s)

	//multi var
	var (
		v1 int
		v2 string
		v3 []int
		v4 bool
	)
	v2 = "你好"
	//fmt.Println(v1, v2, v3)
	fmt.Printf("%v\t%v\t%v\t%v\n",v1,v2,v3,v4)

}
