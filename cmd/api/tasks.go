package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show a task")
}
