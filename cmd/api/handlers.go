package main

import (
	"encoding/json"
	"fmt"
	"go-restapi/inernal/models"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Home Page",
		Version: "1.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-03-24", "2000-01-16")
	harrypotter := models.Movie{
		ID:          1,
		Title:       "Harry Potter",
		Description: "Awesome movie",
		ReleaseDate: rd,
		RunTime:     146,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, harrypotter)
	rd, _ = time.Parse("2006-03-24", "1990-05-26")
	terminator := models.Movie{
		ID:          2,
		Title:       "Terminator",
		Description: "Awesome movie",
		ReleaseDate: rd,
		RunTime:     116,
		MPAARating:  "R",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, terminator)
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
