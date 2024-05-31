package main

import (
	"github.com/gin-gonic/gin"
	"lightcomd/database"
	"lightcomd/endpoints"
)

func main() {
	// Create (if not created) and connect to database
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	// Set up gin
	r := gin.Default()
	r.GET("/publicKey", endpoints.GetServerPublicKey)
	r.PUT("/register", endpoints.Register)
	r.POST("/login", endpoints.Login)
	r.PUT("/send", endpoints.SendMessage)
	r.GET("/new", endpoints.NewMessages)
	r.GET("/newWS", endpoints.NewMessagesWebsocket)
	r.GET("/fetch/:from", endpoints.FetchMessagesFromUser)

	// Run API server
	if err := r.Run(); err != nil {
		panic(err)
	}
}
