package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// ErrorResponse is a error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandler unwraps returned value and if errors occur, responses them to the user
func ErrorHandler[T any](c *gin.Context, f func() (*T, error, error)) *T {
	ret, userError, serverError := f()

	if userError != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{Error: userError.Error()})
		return nil
	} else if serverError != nil {
		log.Println(serverError)
		c.IndentedJSON(http.StatusInternalServerError, &ErrorResponse{Error: "internal error, incident was reported"})
		return nil
	}

	return ret
}

// ErrorHandlerWebsocket does the same as ErrorHandler, but it uses websocket connection
func ErrorHandlerWebsocket[T any](conn *websocket.Conn, f func() (*T, error, error)) *T {
	ret, userError, serverError := f()

	if userError != nil {
		conn.WriteJSON(&ErrorResponse{Error: userError.Error()})
		return nil
	} else if serverError != nil {
		log.Println(serverError)
		conn.WriteJSON(&ErrorResponse{Error: "internal error, incident was reported"})
		return nil
	}

	return ret
}
