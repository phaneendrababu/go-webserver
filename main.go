package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/*func FileServer(root FileSystem) Handler
	FileServer returns a handler that serves HTTP requests with the
	contents of the file system rooted at root.*/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting Server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
func helloHandler(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(w, "404 not found", 404)
	}
	if request.Method != "GET" {
		http.Error(w, "method not supported", 404)
	}

	fmt.Fprintf(w, "hello!!")
}

func formHandler(w http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "not able to process the form")
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(w, "name:%s\n", name)
	fmt.Fprintf(w, "address:%s\n", address)
}
