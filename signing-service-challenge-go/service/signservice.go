package service

import (
	"encoding/base64"
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/mapper"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

const (
	ecc_key = "ECC"
	rsa_key = "RSA"
)

func SaveDevice(deviceRequest domain.DeviceRequest) (domain.CreateDeviceResponse, error) {
	device := mapper.ConvertDeviceRequestToDevice(deviceRequest)

	keyPair := crypto.GenerateKeyPair(device.Algorithm)

	device.KeyPair = keyPair
	deviceOutput := persistence.SaveDevice(device)

	response := mapper.ConvertDeviceOutputToResponse(deviceOutput)

	fmt.Println(deviceOutput)
	return response, nil
}

func GetAllDevices() []domain.CreateDeviceResponse {
	devicesOutput := persistence.GetAllDevices()
	responseList := make([]domain.CreateDeviceResponse, len(devicesOutput))
	counter := 0
	for _, val := range devicesOutput {
		responseList[counter] = mapper.ConvertDeviceOutputToResponse(val)
		counter++
	}
	return responseList
}

func GetDeviceDetails(id string) (domain.CreateDeviceResponse, error) {
	deviceOutput, err := persistence.GetDeviceDetails(id)
	if err != nil {
		return domain.CreateDeviceResponse{}, err
	}
	response := mapper.ConvertDeviceOutputToResponse(deviceOutput)

	return response, nil
}

func SaveSignature(securedToken domain.SignatureRequest) (domain.SignatureResponse, error) {
	requestVO, err := mapper.ConvertStringTokenToSignatureRequestVO(securedToken)
	if err != nil {
		fmt.Println("Test1")
		fmt.Println(err)
		return domain.SignatureResponse{}, err
	}

	device, err := persistence.GetDeviceDetails(requestVO.DeviceId)
	if err != nil {
		fmt.Println("Test2")
		fmt.Println(err)
		return domain.SignatureResponse{}, err
	}

	ok := validateToken(device, requestVO)
	if !ok {
		fmt.Println("Test3")
		fmt.Println(err)
		err = fmt.Errorf("invalid secured token")
		return domain.SignatureResponse{}, err
	}
	newSignature, err := generateSignature(device.Algorithm, []byte(requestVO.Data), device.KeyPair)
	if err != nil {
		fmt.Println("Test4")
		fmt.Println(err)
		return domain.SignatureResponse{}, err
	}
	_, err = persistence.SaveSignature(requestVO.DeviceId, newSignature)
	if err != nil {
		fmt.Println("Test5")
		fmt.Println(err)
		return domain.SignatureResponse{}, err
	}
	response := domain.SignatureResponse{
		Signature:  newSignature,
		SignedData: securedToken.SecuredData,
	}
	return response, nil
}

func generateSignature(algorithm string, dataToBeSigned []byte, encodedPublicKey domain.KeyPair) (string, error) {

	var signer crypto.Signer
	switch algorithm {
	case rsa_key:
		signer = crypto.RSAMarshaler{}
	case ecc_key:
		signer = crypto.ECCMarshaler{}
	}
	response, err := signer.Sign(dataToBeSigned, encodedPublicKey)
	if err != nil {
		return "", err
	}
	encodedSignature := base64.StdEncoding.EncodeToString([]byte(string(response[:])))
	return encodedSignature, nil
}

func validateToken(device domain.Device, requestVO domain.SignatureRequestVO) bool {

	var response bool
	if requestVO.Counter == 0 {
		fmt.Println(requestVO.Counter == device.Counter)
		fmt.Println(requestVO.LastSignature == base64.StdEncoding.EncodeToString([]byte(device.Id)))
		response = requestVO.Counter == device.Counter &&
			requestVO.LastSignature == base64.StdEncoding.EncodeToString([]byte(device.Id))
	} else {
		response = requestVO.Counter == device.Counter &&
			requestVO.LastSignature == device.Signatures[len(device.Signatures)-1]

	}

	return response
}
