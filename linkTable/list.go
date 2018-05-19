package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node
}


func main() {
	node_a := Node{Val: Student{Id: 1, Name: "A"}}
	node_b := Node{Val: Student{Id: 2, Name: "B"}}
	node_c := Node{Val: Student{Id: 3, Name: "C"}}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c

	p:=&node_a
    printNodeList(p)

	fmt.Println("reverse:")
	printNodeList(reverseList(p))

}
