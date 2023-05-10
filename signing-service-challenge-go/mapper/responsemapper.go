package mapper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

const (
	tokenSpearator = "_"
)

func ConvertDeviceRequestToDevice(deviceRequest domain.DeviceRequest) domain.Device {
	device := domain.Device{
		Algorithm:  deviceRequest.Algorithm,
		Counter:    0,
		Signatures: nil,
		Label:      deviceRequest.Label,
	}

	return device
}

func ConvertDeviceOutputToResponse(device domain.Device) domain.CreateDeviceResponse {

	response := domain.CreateDeviceResponse{
		Id:         device.Id,
		Label:      device.Label,
		Counter:    device.Counter,
		Signatures: device.Signatures,
	}

	return response

}

func ConvertStringTokenToSignatureRequestVO(securedToken domain.SignatureRequest) (domain.SignatureRequestVO, error) {

	tokenArr := strings.Split(securedToken.SecuredData, tokenSpearator)
	var response domain.SignatureRequestVO
	if len(tokenArr) == 3 {
		deviceCounter, _ := strconv.Atoi(tokenArr[0])
		response = domain.SignatureRequestVO{
			DeviceId:      securedToken.DeviceId,
			Counter:       deviceCounter,
			Data:          tokenArr[1],
			LastSignature: tokenArr[2],
		}
	} else {
		return response, fmt.Errorf("invalid secured token")
	}

	return response, nil
}
