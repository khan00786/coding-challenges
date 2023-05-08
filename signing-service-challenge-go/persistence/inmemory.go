package persistence

import (
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

var signatureData map[string]domain.Device

func Initialize() {
	signatureData = make(map[string]domain.Device)
}

// *********DEVICE OPERATIONS******************
func SaveDevice(deviceDetails domain.Device) domain.Device {
	signatureData[deviceDetails.Id] = deviceDetails
	return signatureData[deviceDetails.Id]
}

func GetDeviceDetails(deviceId string) (domain.Device, error) {
	device, ok := signatureData[deviceId]
	if ok {
		return device, nil
	} else {
		return domain.Device{}, fmt.Errorf("No Device found for ID: " + deviceId)
	}
}

// *********SIGNATURE OPERATIONS******************
func SaveSignature(deviceId string, signature string) (domain.Device, error) {
	device, err := GetDeviceDetails(deviceId)
	if err != nil {
		return domain.Device{}, err
	}
	device.Signatures = append(device.Signatures, signature)
	return device, nil
}

func GetLastSignature(deviceId string) (string, error) {
	device, ok := signatureData[deviceId]
	if ok {
		lastIndex := device.Counter - 1
		return device.Signatures[lastIndex], nil
	} else {
		return "", fmt.Errorf("No Signature found for ID: " + deviceId)
	}
}

func GetAllDevices() map[string]domain.Device {
	return signatureData
}
