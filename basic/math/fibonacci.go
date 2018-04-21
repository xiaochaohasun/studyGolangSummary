package main

import (
	"flag"
	"fmt"
)

//递归 简洁
func fib(n int) int {
	if n <= 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//迭代算法，速度快
func fib2(n int) int {
	if n<=2{
		return n
	}

	x, y := 1, 2
	for i:=2; i < n; i++ {
		x, y = y, y+x

	}
	return y
}

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "the number of fibonacci")
	flag.Parse()
	var f = fib(num)
	var f2=fib2(num)

	fmt.Printf("fibonacci of %d is: %d\n", num, f)
	fmt.Printf("fibonacci-2 of %d is: %d\n", num, f2)
}
