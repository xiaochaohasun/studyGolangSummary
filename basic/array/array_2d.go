package main

import "fmt"

func main() {
	a := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	a[1][3]=9

	row := len(a)
	col := len(a[0])

	fmt.Println(a, row, col)
}
