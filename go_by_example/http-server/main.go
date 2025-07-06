package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	for i := range 10 {
		fmt.Fprintf(w, "Hi, I am ritesh khadka %d", i)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
