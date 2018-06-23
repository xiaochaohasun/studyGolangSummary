package main

import (
	"fmt"
)

func recoverdall(){
	defer func() {
		if err:=recover();err!=nil{
			//debug.PrintStack()
			fmt.Println(err)
		}
	}()
	call()
}

func call(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	panic("done")
}


func main(){
	recoverdall()
}