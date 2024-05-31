package util

import "golang.org/x/crypto/curve25519"

// PublicFromPrivate returns public key which matches to given private key (uses X25519/Curve25519)
func PublicFromPrivate(privateKey [32]byte) [32]byte {
	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	return publicKey
}

// CalculateSharedSecret creates shared secret calculated from given private and public key (uses X25519/Curve25519)
func CalculateSharedSecret(privateKey, publicKey [32]byte) ([32]byte, error) {
	sharedSecret, err := curve25519.X25519(privateKey[:], publicKey[:])
	if err != nil {
		return [32]byte{}, err
	}

	return [32]byte(sharedSecret), nil
}
