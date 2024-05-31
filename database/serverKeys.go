package database

import (
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"lightcomd/util"
)

var (
	serverPublicKey  [32]byte
	serverPrivateKey [32]byte
)

// KeysModel is a storage model containing server's private and public key
type KeysModel struct {
	gorm.Model
	PublicKey  []byte
	PrivateKey []byte
}

// GenerateKeys generates or reads from database keys for server
func GenerateKeys() error {
	var keysModel KeysModel
	if err := db.First(&keysModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		serverPrivateKey = [32]byte(util.RandomBytes(32))
		serverPublicKey = util.PublicFromPrivate(serverPrivateKey)

		return db.Create(&KeysModel{
			PublicKey:  serverPublicKey[:],
			PrivateKey: serverPrivateKey[:],
		}).Error
	} else if err != nil {
		return err
	}

	serverPublicKey = [32]byte(keysModel.PublicKey)
	serverPrivateKey = [32]byte(keysModel.PrivateKey)

	return nil
}

// GetPublicKey get hex-encoded server's public key
func GetPublicKey() string {
	return hex.EncodeToString(serverPublicKey[:])
}
