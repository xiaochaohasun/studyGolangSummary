package main

import "fmt"

func printNodeList(p *Node){
	for p !=nil{
		node:=*p
		fmt.Printf("id:%d,name:%v,%p\n",node.Val.Id,node.Val.Name,node.Next)
		p=node.Next
	}
}
