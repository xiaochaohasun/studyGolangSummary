package main

import "fmt"

func main(){
	a1 :=[3]int{1,2}
	fmt.Println(a1)

	a2 :=[2]int{1,2}

	fmt.Println(a2)

	//error,type not different
	//fmt.Println(a1 == a2)

	for i:=0;i<3;i++{
		fmt.Printf("a1[%d]: %d, &a1[%d]: %p, &a1: %p\n", i, a1[i], i, &a1[i], &a1)
	}
}
