package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "443123",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "442321",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Jane",
			LastName:  "Hopper",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		ISBN:  "444132",
		Title: "Movie 3",
		Director: &Director{
			FirstName: "Jon",
			LastName:  "Snow",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovies).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
