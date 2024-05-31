package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lightcomd/database"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

// NewMessagesWebsocket is a websocket endpoint which sends to the user message about who sent them messages
func NewMessagesWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, &ErrorResponse{"use websocket"})
		return
	}
	defer conn.Close()

	userId := AuthenticateViaWebsocket(conn)
	if userId == nil {
		return
	}

	newMessages := ErrorHandler(c, func() (*map[string]int, error, error) {
		return database.NewMessages(*userId)
	})
	if newMessages == nil {
		return
	}
	conn.WriteJSON(*newMessages)

	database.OnlineUsers[*userId] = make(chan map[string]int)
	err = nil
	for err == nil {
		newMessages := <-database.OnlineUsers[*userId]
		err = conn.WriteJSON(newMessages)
	}

	log.Println(err)
	delete(database.OnlineUsers, *userId)
}

// NewMessages is an endpoint which responses about who and how many messages sent to requesting user
func NewMessages(c *gin.Context) {
	userId := Authenticate(c)
	if userId == nil {
		return
	}

	newMessages := ErrorHandler(c, func() (*map[string]int, error, error) {
		return database.NewMessages(*userId)
	})
	if newMessages == nil {
		return
	}

	c.IndentedJSON(http.StatusOK, newMessages)
}

// FetchMessagesFromUser is an endpoint which returns all messages sent by specified user to the requesting user
func FetchMessagesFromUser(c *gin.Context) {
	userId := Authenticate(c)
	if userId == nil {
		return
	}

	from := c.Param("from")
	if messages := ErrorHandler(c, func() (*[]database.MessageModel, error, error) {
		return database.FetchMessages(*userId, from)
	}); messages != nil {
		c.IndentedJSON(http.StatusOK, messages)
	}
}
