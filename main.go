package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var userDatabase UserDb

func init() {
	userDatabase = UserDb{
		lock:        sync.Mutex{},
		Users:       make(map[uuid.UUID]bool),
		UserToChats: make(map[uuid.UUID][]*Chat),
	}
}

func main() {
	http.HandleFunc("/", handleEmptyGetReq)
	http.HandleFunc("/newUser", handleNewUser)
	http.HandleFunc("/sendMessage", handleSendMessage)
	http.HandleFunc("/getChats", handleGetChats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
