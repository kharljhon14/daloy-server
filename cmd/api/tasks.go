package main

import (
	"net/http"
	"time"

	"github.com/kharljhon14/daloy-server/internal/data"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// holds the information expected from the request body
	var input struct {
		ProjectID int64  `json:"project_id"`
		OwnerID   int64  `json:"owner_id"`
		AssignID  int64  `json:"assign_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Priority  string `json:"priority"`
		Status    string `json:"status"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": input}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundErrorResponse(w, r)
		return
	}

	task := data.Task{
		ID:        id,
		ProjectID: 1,
		OwnerID:   2,
		AssignID:  2,
		Title:     "Creating the database",
		Content:   "Create the database CONTENT",
		Status:    "open",
		Priority:  "low",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
