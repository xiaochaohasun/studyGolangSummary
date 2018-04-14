package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := " ", " "

	// range return <idx,val> pair
	// key is ignored with '_' for avoiding compile error in Go

	// slice s[m:n], 0<=m<=n<=len(s)

	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}

	fmt.Println(s)

}
