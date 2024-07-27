package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("cotent-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted successfully"})
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "11", Isbn: "458399", Title: "Movie one ", Director: &Director{Firstname: "zahid", Lastname: "khan"}})
	movies = append(movies, Movie{ID: "12", Isbn: "458349", Title: "Movie two ", Director: &Director{Firstname: "wahid", Lastname: "khan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting  server  at port  8010\n")
	log.Fatal(http.ListenAndServe(":8010", r))
	fmt.Println("Hello, World!")
}
