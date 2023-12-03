package handlers

import (
	"net/http"
)

var err error

type errs struct {
	Code int
	Msg  string
}

var (
	errors = map[int]errs{
		404: {
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
		},
		405: {
			http.StatusMethodNotAllowed,
			http.StatusText(http.StatusMethodNotAllowed),
		},
		400: {
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		},
		500: {
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		},
	}
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, myErr errs) {
	w.WriteHeader(myErr.Code)
	err := t.ExecuteTemplate(w, "error.html", myErr)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}
