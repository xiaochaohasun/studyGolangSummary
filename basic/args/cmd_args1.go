package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seq string


	// for type:
	// 1.for loop(the only loop statement in Go):for initialization; condition; post {}

	// 2.while loop: for condition {}

	// 3.infinite loop： for {}



	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}


	fmt.Println(s)

}
