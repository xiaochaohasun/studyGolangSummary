package main

import (
	"fmt"
	"os"
)

func main() {
	s := []rune(os.Args[1])

	fmt.Println("origin string: ", string(s))

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]

	}

	fmt.Println("reverse string: ", string(s))
}
