package database

import (
	"bytes"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"lightcomd/util"
)

// UserModel is a database item containing information about users
type UserModel struct {
	gorm.Model
	UserId string

	PublicKey []byte
}

// NewUser creates user, saves it to the database and returns user id assigned to given public key
func NewUser(publicKey [32]byte) (*string, error, error) {
	userId := GetFreeUserId()

	user := UserModel{
		PublicKey: publicKey[:],
		UserId:    userId,
	}

	return &userId, nil, db.Create(&user).Error
}

// GetFreeUserId return user id which is not used by another user
func GetFreeUserId() string {
	userId := uuid.New().String()

	var user UserModel
	for db.First(&user, "user_id = ?", userId).Error == nil {
		userId = uuid.New().String()
	}

	return userId
}

// Authenticate checks if given shared secret is correct for specified user
func Authenticate(userId string, sharedSecretFromUser [32]byte) bool {
	var user UserModel
	if err := db.First(&user, "user_id = ?", userId).Error; err != nil {
		return false
	}

	sharedSecret, err := util.CalculateSharedSecret(serverPrivateKey, [32]byte(user.PublicKey))
	if err != nil {
		return false
	}

	return bytes.Equal(sharedSecret[:], sharedSecretFromUser[:])
}

// DoesUserExists checks if given user id is present in the database
func DoesUserExists(userId string) bool {
	var user UserModel
	err := db.First(&user, "user_id = ?", userId).Error
	return err == nil
}
