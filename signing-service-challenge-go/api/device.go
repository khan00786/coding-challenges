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
	ListAllDevicesReg = regexp.MustCompile(`^\/api\/v0\/devices[\/]*$`)
	GetDeviceReg      = regexp.MustCompile(`^\/api\/v0\/devices\/([a-z0-9]+)[\/]*$`)
	createDevicesReg  = regexp.MustCompile(`^\/api\/v0\/devices[\/]*$`)
	deviceMutex       = &sync.Mutex{}
)

func (s *Server) Device(response http.ResponseWriter, request *http.Request) {

	switch {

	case request.Method == http.MethodGet && ListAllDevicesReg.Match([]byte(request.URL.Path)):
		ListDevices(response, request)
		return
	case request.Method == http.MethodGet && GetDeviceReg.Match([]byte(request.URL.Path)):
		ListDeviceById(response, request)
		return
	case request.Method == http.MethodPost && createDevicesReg.Match([]byte(request.URL.Path)):
		CreateDevice(response, request)
		return
	default:
		MethodNotAllowedError(response, "Invalid Status Method", request.Method)
		return

	}

}

func ListDevices(response http.ResponseWriter, request *http.Request) {
	log.Printf("list devices invoked")
	devicesOutput := service.GetAllDevices()
	WriteAPIResponse(response, http.StatusOK, devicesOutput)
}

func ListDeviceById(response http.ResponseWriter, request *http.Request) {
	log.Printf("list devices by id invoked")
	matches := GetDeviceReg.FindStringSubmatch(request.URL.Path)
	if len(matches) < 2 {
		InternalServerError(response, "Invalid Device ID", request.URL.Path)
		return
	}
	deviceId := matches[1]
	deviceOutput, err := service.GetDeviceDetails(deviceId)
	if err != nil {
		NoContentError(response, "Failed to Find Device: ", err.Error())
	}
	WriteAPIResponse(response, http.StatusOK, deviceOutput)
}

func CreateDevice(response http.ResponseWriter, request *http.Request) {
	log.Printf("create device invoked")
	decoder := json.NewDecoder(request.Body)
	var device domain.DeviceRequest
	err := decoder.Decode(&device)
	if err != nil {
		InternalServerError(response, "Invalid JSON: ", err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(device)
	if err != nil {
		InternalServerError(response, "Invalid JSON: ", err.Error())
		return
	}
	deviceOutput, err := service.SaveDevice(device, deviceMutex)
	if err != nil {
		InternalServerError(response, "Failed to Save Device: ", err.Error())
		return
	}
	WriteAPIResponse(response, http.StatusCreated, deviceOutput)
}
