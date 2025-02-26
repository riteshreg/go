package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Store short -> long URLs
var urlStore = make(map[string]string)
var mu sync.Mutex // Prevents race conditions

// Generate a random 6-character short code
func generateShortCode() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

// Handle URL shortening
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		LongURL string `json:"long_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()

	mu.Lock()
	urlStore[shortCode] = request.LongURL
	mu.Unlock()

	response := map[string]string{"short_url": "http://localhost:8080/" + shortCode}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handle redirection
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // Extract short code from URL

	mu.Lock()
	longURL, exists := urlStore[shortCode]
	mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
