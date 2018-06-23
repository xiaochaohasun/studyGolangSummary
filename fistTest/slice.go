package main

import "fmt"

func main(){
	//1
	s:=make([]int,3)
	//2
	ss:=make([]int,0)
	s= append(s, 1,2,3)
	ss= append(ss, 1,2,3)
	fmt.Println(s)
	fmt.Println(ss)
}

//1
//0,0,0,1,2,3

//2
//1,2,3