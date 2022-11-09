package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(W http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(W, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(W, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(W, "hello")

}

func formHandler(W http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(W, "ParseForm() error :%v", err)
		return
	}
	fmt.Fprintf(W, "Post request Was successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(W, "Name = %s \n", name)
	fmt.Fprintf(W, "Address = %s", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
