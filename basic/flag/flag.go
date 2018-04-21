package main

import (
	"flag"
	"fmt"
)

func main() {
	//flag.String 返回字符串指针
	var file = flag.String("f", "../point/point.go", "the file name")
	var filen = flag.Bool("n", false, "enter")
	flag.Parse()
	fmt.Println("the file name is:", *file)
	if *filen {
		fmt.Println()
	}

}
