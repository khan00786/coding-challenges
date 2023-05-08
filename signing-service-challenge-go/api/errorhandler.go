package api

import (
	"net/http"
)

func MethodNotAllowedError(response http.ResponseWriter, errorMessage string) {
	WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
		http.StatusText(http.StatusMethodNotAllowed),
	})
}

func InternalServerError(response http.ResponseWriter, errorMessage string) {
	WriteErrorResponse(response, http.StatusInternalServerError, []string{
		http.StatusText(http.StatusInternalServerError),
		errorMessage,
	})
}

func NoContentError(response http.ResponseWriter, errorMessage string) {
	WriteErrorResponse(response, http.StatusNoContent, []string{
		http.StatusText(http.StatusNoContent),
		errorMessage,
	})
}
