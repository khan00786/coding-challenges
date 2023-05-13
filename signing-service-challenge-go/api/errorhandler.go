package api

import (
	"net/http"
)

func MethodNotAllowedError(response http.ResponseWriter, errorMessage string, errorDetails string) {
	WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
		http.StatusText(http.StatusMethodNotAllowed),
		errorMessage,
		errorDetails,
	})
}

func InternalServerError(response http.ResponseWriter, errorMessage string, errorDetails string) {
	WriteErrorResponse(response, http.StatusInternalServerError, []string{
		http.StatusText(http.StatusInternalServerError),
		errorMessage,
		errorDetails,
	})
}

func NoContentError(response http.ResponseWriter, errorMessage string, errorDetails string) {
	WriteErrorResponse(response, http.StatusNoContent, []string{
		http.StatusText(http.StatusNoContent),
		errorMessage,
		errorDetails,
	})
}
