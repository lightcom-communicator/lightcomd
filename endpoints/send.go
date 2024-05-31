package endpoints

import (
	"github.com/gin-gonic/gin"
	"lightcomd/database"
	"net/http"
)

// SendMessage is an endpoint for sending messages
func SendMessage(c *gin.Context) {
	var request database.MessageModel
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"bad request"})
		return
	}

	userId := Authenticate(c)
	if userId == nil {
		return
	}

	if ErrorHandler(c, func() (*struct{}, error, error) {
		return database.SendMessage(*userId, request.ToUser, request.Content)
	}) != nil {
		c.IndentedJSON(http.StatusCreated, &gin.H{"sent": true})
	}
}
