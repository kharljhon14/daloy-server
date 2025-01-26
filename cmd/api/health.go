package main

import "net/http"

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":     "available",
		"enviroment": app.config.env,
		"version":    version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
