package domain

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

// DEVICE Value Objects
type Device struct {
	Id         string   `json:"id"`
	Algorithm  string   `json:"algorithm"`
	Counter    int      `json:"counter"`
	Signatures []string `json:"signatures"`
	Label      string   `json:"label"`
	KeyPair    KeyPair  `json:"keyPair"`
}

type DeviceRequest struct {
	Id        string `json:"id"`
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

type CreateDeviceResponse struct {
	Id         string   `json:"id"`
	Label      string   `json:"label"`
	Counter    int      `json:"counter"`
	Signatures []string `json:"signatures"`
}

// SIGNATURE Value Objects
type SignatureRequest struct {
	DeviceId    string `json:"deviceId"`
	SecuredData string `json:"SecuredData"`
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
