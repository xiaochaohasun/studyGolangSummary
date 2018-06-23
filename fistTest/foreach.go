package main

import (
	"fmt"
)

type gother struct {
	name string
	age  int
}

func main() {
	var m=make(map[string] *gother)
	var l = []gother{
		{name: "tom", age: 20},
		{name: "kevin", age: 28},
		{name: "hasun", age: 30},
	}

	for i, r := range l {
		fmt.Println(r.name, r.age)
		m[r.name] = &l[i]
	}

	for k,v := range m {
		fmt.Printf("key=%s,value=%v\n",k,v)
	}
}
