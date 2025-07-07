package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() %v", err)
		return
	}

	fmt.Fprint(w, "POST Request Successful")

	name := r.Form.Get("name")
	age := r.Form.Get("age")
	gender := r.Form.Get("gender")

	fmt.Fprintf(w, "your name is %s\n", name)
	fmt.Fprintf(w, "your age is %s\n", age)
	fmt.Fprintf(w, "your gender is %s\n", gender)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello Guys!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listing at port 8080")
}
