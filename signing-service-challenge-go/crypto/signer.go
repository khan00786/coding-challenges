package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte, encodedPublicKey domain.KeyPair) ([]byte, error)
}

func (marshaller RSAMarshaler) Sign(dataToBeSigned []byte, encodedPublicKey domain.KeyPair) ([]byte, error) {
	keyPair, err := marshaller.Unmarshal(encodedPublicKey.PrivateKey)

	if err != nil {
		return []byte{}, err
	}

	hash := sha512.New()
	hash.Write(dataToBeSigned)
	digest := hash.Sum(nil)

	encryptedBytes, err := rsa.SignPSS(
		rand.Reader,
		keyPair.Private,
		crypto.SHA512,
		digest,
		nil,
	)

	if err != nil {
		return []byte{}, err
	}
	return encryptedBytes, nil
}

func (marshaller ECCMarshaler) Sign(dataToBeSigned []byte, encodedPublicKey domain.KeyPair) ([]byte, error) {

	keyPair, err := marshaller.Decode(encodedPublicKey.PrivateKey)

	if err != nil {
		return []byte{}, err
	}

	encryptedBytes, err := ecdsa.SignASN1(
		rand.Reader,
		keyPair.Private,
		dataToBeSigned,
	)
	if err != nil {
		return []byte{}, err
	}
	return encryptedBytes, nil
}
