package main

import (
	"fmt"
	"net/http"
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

	fmt.Fprintf(w, "show the details of movie of id: %v", id)
}