package persistence

import (
	"fmt"
	"sync"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

var signatureData map[string]*domain.Device

func Initialize() {
	signatureData = make(map[string]*domain.Device)
}

// *********DEVICE OPERATIONS******************
func SaveDevice(deviceDetails domain.Device, deviceMutex *sync.Mutex) (domain.Device, error) {
	var response domain.Device
	deviceMutex.Lock()
	_, ok := signatureData[deviceDetails.Id]
	if !ok {
		signatureData[deviceDetails.Id] = &deviceDetails
		response = *signatureData[deviceDetails.Id]
		deviceMutex.Unlock()
	} else {
		deviceMutex.Unlock()
		return response, fmt.Errorf("Device already exist: " + deviceDetails.Id)
	}
	return response, nil
}

func GetDeviceDetails(deviceId string) (*domain.Device, error) {
	device, ok := signatureData[deviceId]
	if ok {
		return device, nil
	} else {
		return &domain.Device{}, fmt.Errorf("No Device found for ID: " + deviceId)
	}
}

func GetAllDevices() map[string]*domain.Device {
	return signatureData
}

// *********SIGNATURE OPERATIONS******************
func SaveSignature(device *domain.Device, signature string) (domain.Device, error) {

	device.Signatures = append(device.Signatures, signature)
	device.Counter++
	return *device, nil
}
