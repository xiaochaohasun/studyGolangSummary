package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	seen := make(map[string]bool)
	reader:= bufio.NewReader(os.Stdin)
	for{
		fmt.Print("> ")
		line,_:=reader.ReadString('\n')
		line=strings.TrimSuffix(line,"\n")
		if !seen[line]{
			seen[line]=true
			fmt.Println(line)
		}
	}
}
