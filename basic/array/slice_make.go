package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 2)
	printSlice("b", b)

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	var s []int
	s = append(a, 1)
	fmt.Println(s)

	s = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(s)

	s2 := b[:]
	for i := 0; i < 5; i++ {
		s2 = append(s2, i)
		fmt.Printf("cap: %d, len:%d, s:%p, value:%v\n", cap(s2), len(s2), s2, s2)

	}
	//fmt.Printf("cap: %d, len:%d, s:%p, value:%v\n",cap(s2),len(s2),s,s)

	uniq := []string{"kevin", "kevin", "kevin", "hasun", "neal"}
	dup := Rm_duplicate(uniq)

	fmt.Printf("dup: %v\n", dup)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func Rm_duplicate(list []string) []string {
	var x []string = []string{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
