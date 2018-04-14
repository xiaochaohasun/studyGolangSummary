package main

import (
	"fmt"
	"math"
)

func main() {
	a := -2.0                // have .0  default float type
	fmt.Println(math.Abs(a)) //result:2
	b := math.Pow(a, 2)
	fmt.Println(b) //result:4

}
