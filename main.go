package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

var userDatabase UserDb

func init() {
	userDatabase = UserDb{
		Users:         make(map[uuid.UUID]bool),
		UsersToUsers:  make(map[uuid.UUID]uuid.UUID),
		UserOneToChat: make(map[uuid.UUID]*Chat),
	}
}

func main() {
	http.HandleFunc("/", handleEmptyGetReq)
	http.HandleFunc("/newUser", handleNewUser)
	http.HandleFunc("/sendMessage", handleSendMessage)
	http.HandleFunc("/getChats", handleGetChats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
