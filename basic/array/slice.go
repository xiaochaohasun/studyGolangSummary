package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]
	s1 := s[:4]
	fmt.Println(s, s1)

	//string
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "kevin"

	fmt.Println(a, b)
	fmt.Println(names)

	//sliece len and cap(容量)

	array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j'}
	sliece_a := array_a[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(sliece_a), cap(sliece_a), sliece_a)

	sliece_b := sliece_a[0:8]
	fmt.Printf("len=%d cap=%d %v\n", len(sliece_b), cap(sliece_b), sliece_b)




}
