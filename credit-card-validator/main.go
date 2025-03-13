package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Luhn(card_number string) bool {

	var revert bool = false
	var sum int = 0

	for i := len(card_number) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(card_number[i]))
		if err != nil {
			return false
		}

		if revert {
			num = num * 2
			if num > 9 {
				num = num - 9
			}
		}

		sum += num
		revert = !revert
	}

	return sum%10 == 0

}

type PageData struct {
	Title   string
	Heading string
	Content string
}

func handler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Go HTML Templates",
		Heading: "Welcome to Go Templates",
		Content: "This is an example of using HTML templates in Go.",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "The Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the card number from the form
	cardNumber := r.FormValue("card_number")

	if cardNumber == "" {
		http.Error(w, "Card number is missing", http.StatusBadRequest)
		return
	}

	isCardValid := Luhn(cardNumber)

	if isCardValid {
		fmt.Fprintf(w, "the card %s is valid", cardNumber)
		return
	}

	fmt.Fprintf(w, "%s is not valid card", cardNumber)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/validate-card", cardHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
