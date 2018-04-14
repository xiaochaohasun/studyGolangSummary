package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hello world!")
}

func main()  {
	log.Println("Starting http server on localhost:8000")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8000",nil)

}