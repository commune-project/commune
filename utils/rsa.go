package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const keyBitSize = 2048

// GenerateRsaKeys returns a public key and a private key, both in PEM format.
func GenerateRsaKeys() (pubPEM []byte, privPEM []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keyBitSize)
	if err != nil {
		return nil, nil
	}

	publicKey := privateKey.PublicKey

	pubBytes := pem.EncodeToMemory(
		&pem.Block{
			Bytes: x509.MarshalPKCS1PublicKey(&publicKey),
			Type:  "PUBLIC KEY",
		})
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			Type:  "RSA PRIVATE KEY",
		})
	return pubBytes, privBytes
}

// ParsePublicKey converts PEM into crypto.PublicKey
func ParsePublicKey(pubPEM []byte) (crypto.PublicKey, error) {
	p, _ := pem.Decode(pubPEM)
	if p != nil {
		return nil, errors.New("unable to decode key")
	}
	return x509.ParsePKCS1PublicKey(p.Bytes)
}
