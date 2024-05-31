package database

import (
	"errors"
	"gorm.io/gorm"
)

// MessageModel is a database object containing information about message
type MessageModel struct {
	gorm.Model
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Content  string `json:"content"` // encrypted and encoded in hex
}

var OnlineUsers = make(map[string]chan map[string]int)

// SendMessage saves message to the database
func SendMessage(from, to, content string) (*struct{}, error, error) {
	message := MessageModel{
		FromUser: from,
		ToUser:   to,
		Content:  content,
	}

	if DoesUserExists(message.ToUser) {
		serverError := db.Create(&message).Error
		if serverError != nil {
			return nil, nil, serverError
		}

		if channel, ok := OnlineUsers[message.ToUser]; ok {
			channel <- map[string]int{
				message.FromUser: 1,
			}
		}

		return &struct{}{}, nil, nil
	}

	return nil, errors.New("user not found"), nil
}

// NewMessages returns numbers of message sent to specified user by other users
func NewMessages(userId string) (*map[string]int, error, error) {
	var results []struct {
		FromUser string
		Amount   int
	}

	if err := db.
		Model(&MessageModel{}).
		Select("from_user, count(*) as amount").
		Where("to_user = ?", userId).
		Group("from_user").
		Scan(&results).Error; err != nil {
		return nil, errors.New("user not found"), nil
	}

	list := make(map[string]int)
	for _, result := range results {
		list[result.FromUser] = result.Amount
	}

	return &list, nil, nil
}

// FetchMessages get all messages sent
func FetchMessages(userId, from string) (*[]MessageModel, error, error) {
	var messagesList []MessageModel
	if err := db.
		Find(&messagesList, "from_user = ? AND to_user = ?", from, userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found"), nil
	} else if err != nil {
		return nil, nil, err
	}

	db.Delete(&MessageModel{}, "from_user = ? AND to_user = ?", from, userId)
	return &messagesList, nil, nil
}
