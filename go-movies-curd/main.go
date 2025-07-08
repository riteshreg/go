package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi, Home")
}

// fetch list of movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get a movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting params from /movie/id
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

// insert a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// creating a emply movie variable
	var movie Movie

	// decoding and assing the json to movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// creating and id for movie
	movie.ID = strconv.Itoa(rand.Int())

	// finally appending
	movies = append(movies, movie)

	// also returning the created movie
	json.NewEncoder(w).Encode(movie)
}

// delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting params /movie/id from mux vars
	params := mux.Vars(r)

	for i, movie := range movies {

		if movie.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// first get id
	params := mux.Vars(r)
	// get json and store that movie in variable

	// delect that the movie
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)

			break
		}
	}
	// and store new movie

}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "np_1", Title: "Test", Director: &Director{Firstname: "Ritesh", Lastname: "Khadka"},
	})

	movies = append(movies, Movie{
		ID: "2", Isbn: "np_2", Title: "Best", Director: &Director{Firstname: "Nayan", Lastname: "Regmi"},
	})

	r.HandleFunc("/", HomeHandler)

	// movies route
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listing at port 8000")
	log.Fatal(srv.ListenAndServe())
}
