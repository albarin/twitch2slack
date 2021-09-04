package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) badRequestResponse(w http.ResponseWriter, err error) {
	app.errorResponse(w, http.StatusBadRequest, err.Error())
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Err(err).
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Send()

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, http.StatusInternalServerError, message)
}

func (app *application) errorResponse(w http.ResponseWriter, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
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
