package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	paramId := r.PathValue("id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.Marshal(data)
	// js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	// Set the max request body size to 1mb
	max_bytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_bytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the request body
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		// Check whether the error has the type of *jsonSyntaxError
		// Check if a valid JSON
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")
		// Check whether a field has incorrect JSON type for the target dst
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field (%s)", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		// Check whether the request body is empty
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		// Override error message from DissalowUnknownFields
		case strings.Contains(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		// Check the size of the request body
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must be not larger than %d bytes", max_bytes)
		// Check if we pass a non-nil pointer to Decode
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	// Check if the request body contains only a single JSON value
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
