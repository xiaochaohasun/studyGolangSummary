package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	//conn.Close()
}

func main() {
	log.Println("start to listen on tcp port 9000")

	l, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal(err)
		print("a")
	}

	for {
		conn, err := l.Accept()
		if err != nil{
			log.Fatal(err)
		}


		go handler(conn)
	}

}
