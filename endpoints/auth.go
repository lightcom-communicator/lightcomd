package endpoints

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lightcomd/database"
	"net/http"
)

// Authenticate checks if access token given by user is assigned to one of the users
func Authenticate(c *gin.Context) *string {
	accessToken := c.GetHeader("Authorization")

	return ErrorHandler(c, func() (*string, error, error) {
		return database.GetUserIdByAccessToken(accessToken)
	})
}

// AuthenticateViaWebsocket does the same as Authenticate but uses websocket connection
func AuthenticateViaWebsocket(conn *websocket.Conn) *string {
	type AuthRequest struct {
		AccessToken string `json:"accessToken"`
	}

	var request AuthRequest
	if err := conn.ReadJSON(&request); err != nil {
		conn.WriteJSON(&ErrorResponse{"unauthorized"})
		return nil
	}

	return ErrorHandlerWebsocket(conn, func() (*string, error, error) {
		return database.GetUserIdByAccessToken(request.AccessToken)
	})
}

// GetServerPublicKey responses server's public key to the requester
func GetServerPublicKey(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, &gin.H{
		"publicKey": database.GetPublicKey(),
	})
}

// Login generates access token and assigns it to the specified user if given shared secret is correct
func Login(c *gin.Context) {
	type LoginRequest struct {
		UserId       string `json:"userId"`
		SharedSecret string `json:"sharedSecret"`
	}

	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"bad request"})
		return
	}

	sharedSecret, err := hex.DecodeString(request.SharedSecret)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"invalid credentials"})
		return
	}

	if database.Authenticate(request.UserId, [32]byte(sharedSecret)) {
		accessToken := ErrorHandler(c, func() (*database.AccessToken, error, error) {
			return database.NewAccessToken(request.UserId)
		})

		if accessToken == nil {
			return
		}

		c.IndentedJSON(http.StatusOK, &accessToken)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, &ErrorResponse{"invalid credentials"})
	}
}

// Register creates a new user
func Register(c *gin.Context) {
	type RegisterRequest struct {
		PublicKey string `json:"publicKey"`
	}

	type RegisterResponse struct {
		UserId string `json:"userId"`
	}

	var request RegisterRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"bad request"})
		return
	}

	publicKey, err := hex.DecodeString(request.PublicKey)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"invalid public key"})
		return
	}

	userId := ErrorHandler(c, func() (*string, error, error) {
		return database.NewUser([32]byte(publicKey))
	})

	if userId != nil {
		c.IndentedJSON(http.StatusCreated, &RegisterResponse{*userId})
	}
}
