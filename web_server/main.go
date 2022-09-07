package main

import(
	"fmt"
	"log"
	"net/http"
)

// ResponseWriter: This is the mechanism used for sending responses to any connected HTTP clients. It's also how response headers are set.
// http.Request: is a pointer to an http.Request. It's how data is retrieved from the web request. For example, the details from a form submission can be accessed through the request pointer.
// ref: https://www.honeybadger.io/blog/go-web-services/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}