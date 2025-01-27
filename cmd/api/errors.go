package main

import (
	"net/http"
)

// Print out the error
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Build the error response
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"message": message}

	if err := app.writeJSON(w, status, env, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// func (app *application) methodNotAllowedErrorResponse(w http.ResponseWriter, r *http.Request) {
// 	message := fmt.Sprintf("the %s method is not supported for this route", r.Method)
// 	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
// }
