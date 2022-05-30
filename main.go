package main

import (
    "net/http" //package for http based web programs
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Inside handler")
    fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
    http.HandleFunc("/", handler) // redirect all urls to the handler function
	fmt.Println("http://localhost:8080")
    http.ListenAndServe("localhost:8080", nil) // listen for connections at port 9999 on the local machine
}