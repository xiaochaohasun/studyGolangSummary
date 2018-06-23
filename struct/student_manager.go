package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Scan(line, &cmd)
		switch cmd {
		case "add":
			var id int
			var name string
			fmt.Scan(line, &cmd, &id, &name)
			stus[name] = Student{Id: id, Name: name}
		case "list":
			for _, v := range stus {
				fmt.Printf("id:%d,name=%v\n", v.Id, v.Name)
			}

		}
	}
}

func add(id int,name string){
	if _,ok:=students[name]; ok{
		fmt.Printf("student %s has already extisted!",name)
	} else{
		students[name]=Student{Id:id,Name:name}
	}

}
