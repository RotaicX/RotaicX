package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func NewKey4096() (*rsa.PrivateKey, error) {
	PrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	return PrivateKey, nil
}

func GetPublicKey(PrivateKey *rsa.PrivateKey) rsa.PublicKey {
	return PrivateKey.PublicKey
}

func PrivateKeyToBase64PrivateKey(PrivateKey *rsa.PrivateKey) string {
	x509Key := x509.MarshalPKCS1PrivateKey(PrivateKey)
	return base64.StdEncoding.EncodeToString(x509Key)
}

func PublicKeyToBase64PublicKey(PublicKey *rsa.PublicKey) string {
	x509Key := x509.MarshalPKCS1PublicKey(PublicKey)
	return base64.StdEncoding.EncodeToString(x509Key)
}

func Base64PrivateKeyToPrivateKey(Base64PrivateKey string) (*rsa.PrivateKey, error) {
	X509PrivateKey, err := base64.StdEncoding.DecodeString(Base64PrivateKey)
	if err != nil {
		return nil, err
	}

	PrivateKey, err := x509.ParsePKCS1PrivateKey(X509PrivateKey)
	if err != nil {
		return nil, err
	}

	return PrivateKey, nil
}

func Base64PublicKeyToPublicKey(Base64PublicKey string) (*rsa.PublicKey, error) {
	X509PublicKey, err := base64.StdEncoding.DecodeString(Base64PublicKey)
	if err != nil {
		return nil, err
	}

	PublicKey, err := x509.ParsePKCS1PublicKey(X509PublicKey)
	if err != nil {
		return nil, err
	}

	return PublicKey, nil
}

func Encrypt(Data string, PublicKey *rsa.PublicKey) ([]byte, error) {
	EncryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		PublicKey,
		[]byte(Data),
		nil)
	if err != nil {
		return nil, err
	}
	return EncryptedBytes, nil
}

func Decrypt(Data string, PrivateKey *rsa.PrivateKey) ([]byte, error) {
	DecryptBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		PrivateKey,
		[]byte(Data),
		nil)
	if err != nil {
		return nil, err
	}
	return DecryptBytes, nil
}
