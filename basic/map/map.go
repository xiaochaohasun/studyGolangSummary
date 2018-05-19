package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["tom"] = 28
	ages["kevin"] = 26
	for k, v := range ages {
		fmt.Printf("name:%v,age=%d\n", k, v)
	}


	ages1:=map[string]int{"xiaochao":25}
	for k, v := range ages1 {
		fmt.Printf("name:%v,age=%d\n", k, v)
	}


	delete(ages,"tom")
	v:=ages["kevin"]
	fmt.Printf("kevin's age=%d\n",v)
}
