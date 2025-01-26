package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /api/v1/health", app.healthHandler)

	// Tasks
	mux.HandleFunc("GET /api/v1/tasks/{id}", app.showTaskHandler)
	mux.HandleFunc("POST /api/v1/tasks", app.createTaskHandler)

	return mux
}
