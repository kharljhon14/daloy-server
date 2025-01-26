package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kharljhon14/daloy-server/internal/data"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
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
		app.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
