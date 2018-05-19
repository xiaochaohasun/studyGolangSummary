package main

import "fmt"

type student struct {
	Id   int
	Name string
	Sex  bool
}

func main() {
	var s student
	s.Id = 1
	s.Name = "xiaochao"

	p:=&s
	p.Name="kevin"



	fmt.Printf("%v\n%v\n",s,*p)
}
