package wg

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

func GenerateKeyPair() (privateKey, publicKey string, err error) {
	var priv [32]byte
	_, err = rand.Read(priv[:])
	if err != nil {
		return "", "", err
	}

	priv[0] &= 248
	priv[31] = (priv[31] & 127) | 64

	var pub [32]byte
	curve25519.ScalarBaseMult(&pub, &priv)

	privateKey = base64.StdEncoding.EncodeToString(priv[:])
	publicKey = base64.StdEncoding.EncodeToString(pub[:])
	return
}
