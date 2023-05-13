package domain

// DEVICE Structs
type Device struct {
	Id         string   `json:"id"`
	Algorithm  string   `json:"algorithm"`
	Counter    int      `json:"counter"`
	Signatures []string `json:"signatures"`
	Label      string   `json:"label"`
	KeyPair    KeyPair  `json:"keyPair"`
}

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

type DeviceRequest struct {
	Algorithm string `json:"algorithm" validate:"required"`
	Label     string `json:"label" validate:"required"`
}

type CreateDeviceResponse struct {
	Id         string   `json:"id"`
	Label      string   `json:"label"`
	Counter    int      `json:"counter"`
	Signatures []string `json:"signatures"`
}

// SIGNATURE Structs
type SignatureRequest struct {
	DeviceId    string `json:"deviceId" validate:"required"`
	SecuredData string `json:"SecuredData" validate:"required"`
}

type SignatureRequestVO struct {
	DeviceId      string `json:"deviceId"`
	Counter       int    `json:"SecuredData"`
	Data          string `json:"data"`
	LastSignature string `json:"lastSignature"`
}

type SignatureResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signedData"`
}

//HEALTH Structs
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
