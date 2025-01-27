package main

import (
	"net/http"
	"time"

	"github.com/kharljhon14/daloy-server/internal/data"
	"github.com/kharljhon14/daloy-server/internal/validator"
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
		app.badRequestErrorRespons(w, r, err)
		return
	}

	v := validator.New()

	// Check project_id
	v.Check(input.ProjectID != 0, "project_id", "Project ID must be provided")
	v.Check(input.ProjectID > 0, "project_id", "Project ID must be a positive integer")

	// Check owner_id
	v.Check(input.OwnerID != 0, "owner_id", "Owner ID must be provided")
	v.Check(input.OwnerID > 0, "owner_id", "Owner ID must be a positive integer")

	// Check assign_id
	v.Check(input.AssignID != 0, "assign_id", "Assign ID must be provided")
	v.Check(input.AssignID > 0, "assign_id", "Assign ID must be a positive integer")

	// Check title
	v.Check(input.Title != "", "title", "Title must be provided")
	v.Check(len(input.Title) <= 500, "title", "Title must not be more thatn 500 bytes long")

	// Check content
	v.Check(input.Content != "", "content", "Content must be provided")
	v.Check(len(input.Content) <= 5000, "content", "Content must not be more thatn 5000 bytes long")

	// Check priority
	v.Check(input.Priority != "", "priority", "Priority must be provided")
	v.Check(v.In(input.Priority, data.AllowedPriorities...), "priority", "Priority should only be (low, medium, high)")

	// Check status
	v.Check(input.Status != "", "status", " must be provided")
	v.Check(v.In(input.Status, data.AllowedStatus...), "status", "Status should only be (done, open, in progress, in queue)")

	if !v.Valid() {
		app.failedValidationErrorResponse(w, r, v.Errors)
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
