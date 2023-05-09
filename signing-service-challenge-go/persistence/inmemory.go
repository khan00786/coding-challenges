package persistence

import (
	"fmt"
	"log"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

var signatureData map[string]*domain.Device

func Initialize() {
	signatureData = make(map[string]*domain.Device)
}

// *********DEVICE OPERATIONS******************
func SaveDevice(deviceDetails domain.Device) domain.Device {
	signatureData[deviceDetails.Id] = &deviceDetails
	log.Printf("SaveDevice %v\n", signatureData)
	return *signatureData[deviceDetails.Id]
}

func GetDeviceDetails(deviceId string) (*domain.Device, error) {
	log.Printf("GetDeviceDetails %v\n", signatureData)
	device, ok := signatureData[deviceId]
	if ok {
		return device, nil
	} else {
		return &domain.Device{}, fmt.Errorf("No Device found for ID: " + deviceId)
	}
}

// *********SIGNATURE OPERATIONS******************
func SaveSignature(deviceId string, signature string) (domain.Device, error) {
	device, err := GetDeviceDetails(deviceId)
	if err != nil {
		return domain.Device{}, err
	}

	log.Printf("SaveSignature1 %v\n", device.Signatures)
	log.Printf("SaveSignature2 %v\n", signature)
	device.Signatures = append(device.Signatures, signature)
	device.Counter++
	log.Printf("SaveSignature3 %v\n", device.Signatures)
	log.Printf("SaveSignature %v\n", signatureData)
	// signatureData[device.Id] = deviceDetails
	return *device, nil
}

func GetLastSignature(deviceId string) (string, error) {
	log.Printf("GetLastSignature %v\n", signatureData)
	device, err := GetDeviceDetails(deviceId)
	if err != nil {
		return "", fmt.Errorf("No Signature found for ID: " + deviceId)
	} else {
		lastIndex := device.Counter - 1
		return device.Signatures[lastIndex], nil
	}
}

func GetAllDevices() map[string]*domain.Device {
	log.Printf("GetAllDevices %v\n", signatureData)
	return signatureData
}
