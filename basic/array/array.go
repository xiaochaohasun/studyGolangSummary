package main

import "fmt"

func main() {

	//1
	//var array [3]int
	//2
	array := [...]int{0, 1, 2}
	fmt.Println(array[0])
	fmt.Println(array[len(array)-1])

	for i, v := range array {
		fmt.Printf("%d %d\n", i, v)
	}

	for i := range array {
		fmt.Printf("%d\n", i)
	}

	for _, v := range array {
		fmt.Printf("%d\n", v)
	}
}
