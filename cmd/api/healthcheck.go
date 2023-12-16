package main

import (
	
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

	// data := map[string]string {
	// 	"status": "available",
	// 	"environment": app.config.env,
	// 	"version": version,
	// }
	env := envelope{
		"status": "status",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered an error and couldn't process your request", http.StatusInternalServerError)
		return
	}
}