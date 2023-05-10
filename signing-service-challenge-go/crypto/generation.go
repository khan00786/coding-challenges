package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

const (
	ecc_key = "ECC"
	rsa_key = "RSA"
)

// RSAGenerator generates a RSA key pair.
type RSAGenerator struct{}

// Generate generates a new RSAKeyPair.
func (g *RSAGenerator) Generate() (*RSAKeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &RSAKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

// ECCGenerator generates an ECC key pair.
type ECCGenerator struct{}

// Generate generates a new ECCKeyPair.
func (g *ECCGenerator) Generate() (*ECCKeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &ECCKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

func GenerateKeyPair(algorithm string) domain.KeyPair {

	switch algorithm {
	case rsa_key:
		return GenerateRSAKeyPair()
	case ecc_key:
		return GenerateECCKeyPair()
	}

	return domain.KeyPair{}
}

func GenerateUUID() string {
	id := uuid.New()
	return fmt.Sprintf("%x", id)
}
