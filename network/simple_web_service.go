package network

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Println("init network simple_web_service")
}

func SimpleWebService() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/person/", personHandle)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal("ListenAndService:", err.Error())
	}
}

func rootHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handle /")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func personHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handle /person")
	fmt.Fprintf(w, "person:"+req.URL.Path[len("/person/"):])
}
