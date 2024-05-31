package database

import (
	"errors"
	"gorm.io/gorm"
	"lightcomd/util"
	"time"
)

// AccessTokenModel is a database item containing information about access tokens
type AccessTokenModel struct {
	gorm.Model
	AccessToken string
	ForUser     string
	ValidUntil  time.Time
}

// AccessToken is a return value from NewAccessToken
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ValidUntil  int64  `json:"validUntil"`
}

// NewAccessToken creates an access token for a specific user
func NewAccessToken(userId string) (*AccessToken, error, error) {
	token := util.RandomBytesHexEncoded(256)

	var model AccessTokenModel
	for db.First(&model, "access_token = ?", token).Error == nil {
		token = util.RandomBytesHexEncoded(256)
	}

	model = AccessTokenModel{
		AccessToken: token,
		ForUser:     userId,
		ValidUntil:  time.Now().AddDate(1, 0, 0),
	}

	return &AccessToken{token, model.ValidUntil.Unix()}, nil, db.Create(&model).Error
}

// GetUserIdByAccessToken returns user id assigned to the given access token
func GetUserIdByAccessToken(accessToken string) (*string, error, error) {
	var model AccessTokenModel
	if err := db.First(&model, "access_token = ?", accessToken).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("invalid access token"), nil
	} else if err != nil {
		return nil, nil, err
	}

	if time.Now().After(model.ValidUntil) {
		db.Delete(&model)
		return nil, errors.New("invalid access token"), nil
	}

	return &model.ForUser, nil, nil
}
