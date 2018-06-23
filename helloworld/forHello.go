package main

import "fmt"

func main() {
	s := "hello 大家好"
	r:=[]rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
	}
}
