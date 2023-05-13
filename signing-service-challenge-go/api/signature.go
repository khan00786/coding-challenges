package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
	"github.com/go-playground/validator/v10"
)

var (
	createSignatureReg = regexp.MustCompile(`^\/api\/v0\/signatures[\/]*$`)
	signatureMutex     = &sync.Mutex{}
)

func (s *Server) Signature(response http.ResponseWriter, request *http.Request) {

	switch {

	case request.Method == http.MethodPost && createSignatureReg.Match([]byte(request.URL.Path)):
		SignTransaction(response, request)
		return
	default:
		MethodNotAllowedError(response, "Invalid Status Method", request.Method)
		return

	}

}

func SignTransaction(response http.ResponseWriter, request *http.Request) {
	log.Printf("sign transaction invoked")
	decoder := json.NewDecoder(request.Body)
	var signRequest domain.SignatureRequest
	err := decoder.Decode(&signRequest)
	if err != nil {
		InternalServerError(response, "Invalid JSON: ", err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(signRequest)
	if err != nil {
		InternalServerError(response, "Invalid JSON: ", err.Error())
		return
	}
	signatureResponse, err := service.SaveSignature(signRequest, signatureMutex)
	if err != nil {
		InternalServerError(response, "Failed to Sign Transaction: ", err.Error())
		return
	}
	WriteAPIResponse(response, http.StatusCreated, signatureResponse)
}
