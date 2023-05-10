package service

import (
	"encoding/base64"
	"fmt"
	"log"

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

	device.KeyPair = crypto.GenerateKeyPair(device.Algorithm)
	device.Id = crypto.GenerateUUID()

	deviceOutput := persistence.SaveDevice(device)

	response := mapper.ConvertDeviceOutputToResponse(deviceOutput)

	log.Printf("device saved successfully:[id] %v\n", device.Id)
	return response, nil
}

func GetAllDevices() []domain.CreateDeviceResponse {
	devicesOutput := persistence.GetAllDevices()
	responseList := make([]domain.CreateDeviceResponse, len(devicesOutput))
	counter := 0
	for _, val := range devicesOutput {
		responseList[counter] = mapper.ConvertDeviceOutputToResponse(*val)
		counter++
	}
	log.Printf("all devices returned successfully:[total count] %v\n", len(devicesOutput))
	return responseList
}

func GetDeviceDetails(id string) (domain.CreateDeviceResponse, error) {
	var response domain.CreateDeviceResponse
	deviceOutput, err := persistence.GetDeviceDetails(id)
	if err != nil {
		log.Println(err)
		return response, err
	}
	response = mapper.ConvertDeviceOutputToResponse(*deviceOutput)
	log.Printf("device returned successfully:[id] %v\n", id)
	return response, nil
}

func SaveSignature(securedToken domain.SignatureRequest) (domain.SignatureResponse, error) {
	var response domain.SignatureResponse
	requestVO, err := mapper.ConvertStringTokenToSignatureRequestVO(securedToken)
	if err != nil {
		log.Println(err)
		return response, err
	}

	device, err := persistence.GetDeviceDetails(requestVO.DeviceId)
	if err != nil {
		log.Println(err)
		return response, err
	}

	ok := validateToken(*device, requestVO)
	if !ok {
		err = fmt.Errorf("invalid secured token")
		log.Println(err)
		return response, err
	}
	newSignature, err := generateSignature(device.Algorithm, []byte(requestVO.Data), device.KeyPair)
	if err != nil {
		log.Println(err)
		return response, err
	}
	_, err = persistence.SaveSignature(requestVO.DeviceId, newSignature)
	if err != nil {
		log.Println(err)
		return domain.SignatureResponse{}, err
	}
	response = domain.SignatureResponse{
		Signature:  newSignature,
		SignedData: securedToken.SecuredData,
	}
	log.Printf("signature saved successfully:[id] %v\n", requestVO.DeviceId)
	return response, nil
}

func generateSignature(algorithm string, dataToBeSigned []byte, encodedPublicKey domain.KeyPair) (string, error) {
	var response string
	var signer crypto.Signer
	switch algorithm {
	case rsa_key:
		signer = crypto.RSAMarshaler{}
	case ecc_key:
		signer = crypto.ECCMarshaler{}
	}
	signature, err := signer.Sign(dataToBeSigned, encodedPublicKey)
	if err != nil {
		return response, err
	}
	response = base64.StdEncoding.EncodeToString(signature)
	return response, nil
}

func validateToken(device domain.Device, requestVO domain.SignatureRequestVO) bool {

	var response bool
	log.Printf("Device Counter: %v\n", requestVO.Counter == device.Counter)
	if device.Counter == 0 {
		log.Printf("Last Signature for 0: %v\n", requestVO.LastSignature == base64.StdEncoding.EncodeToString([]byte(device.Id)))

		response = requestVO.Counter == device.Counter &&
			requestVO.LastSignature == base64.StdEncoding.EncodeToString([]byte(device.Id))
	} else {
		log.Printf("Last Signature for >0: %v\n", device.Signatures[len(device.Signatures)-1] == base64.StdEncoding.EncodeToString([]byte(device.Id)))
		response = requestVO.Counter == device.Counter &&
			requestVO.LastSignature == device.Signatures[len(device.Signatures)-1]

	}

	return response
}
