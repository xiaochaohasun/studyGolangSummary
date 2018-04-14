package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.1f", c)  //result:5.0%
	fmt.Printf("%.1f\n", c) //result:5.0  - only  mac os
}
