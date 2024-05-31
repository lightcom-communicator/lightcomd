package util

import (
	"encoding/hex"
	"math/rand"
)

// RandomBytes creates random bytes array
func RandomBytes(length int) []byte {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = byte(rand.Intn(256))
	}

	return bytes
}

// RandomBytesHexEncoded creates length / 2 bytes array and encodes it in hex
func RandomBytesHexEncoded(length int) string {
	return hex.EncodeToString(RandomBytes(length / 2))
}
