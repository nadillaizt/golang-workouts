package main

import "fmt"
import "net/http"

// handler
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// for simple
	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
	    fmt.Println(err.Error())
	}

	// // more flexible u can set additional fields such as readtimeout etc
	// var address = ":9000"
	// fmt.Printf("server started at %s\n", address)

	// server := new(http.Server)
	// server.Addr = address
	// err := server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	
}
