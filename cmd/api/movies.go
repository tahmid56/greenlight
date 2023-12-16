package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tahmid56/greenlight/internal/data.go"
)

func (app *application) createMovieHanlder(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "creating movie")
}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request){
	
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	movie := data.Movie {
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casablanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server envountered a problem and could not process your request", http.StatusInternalServerError)
	}
}