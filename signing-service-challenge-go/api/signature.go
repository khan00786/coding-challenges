package api

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
)

var (
	createSignatureReg = regexp.MustCompile(`^\/api\/v0\/signatures[\/]*$`)
)

func (s *Server) Signature(response http.ResponseWriter, request *http.Request) {

	switch {

	case request.Method == http.MethodPost && createSignatureReg.Match([]byte(request.URL.Path)):
		SignTransaction(response, request)
		return
	default:
		MethodNotAllowedError(response, "Invalid Status Method")
		return

	}

}

func SignTransaction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var signRequest domain.SignatureRequest
	err := decoder.Decode(&signRequest)
	if err != nil {
		InternalServerError(response, "Invalid JSON")
	}
	signatureResponse, err := service.SaveSignature(signRequest)
	if err != nil {
		InternalServerError(response, "Failed to Sign Transaction")
	}
	WriteAPIResponse(response, http.StatusCreated, signatureResponse)
}
