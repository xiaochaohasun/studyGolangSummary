package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func pringFile(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		//1.fmt.Println(err)
		//2.log.Fatal(err)
		log.Fatal(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	pringFile(os.Args[1])
}
