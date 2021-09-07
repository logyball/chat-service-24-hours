package main

import (
	"log"
	"net/http"
	"sync"
)

var userDatabase UserDb

func init() {
	userDatabase = UserDb{
		lock:        sync.Mutex{},
		curId:       0,
		Users:       map[int32]bool{},
		UserToChats: map[int32][]*Chat{},
	}
}

func main() {
	http.HandleFunc("/", handleEmptyGetReq)
	http.HandleFunc("/newUser", handleNewUser)
	http.HandleFunc("/sendMessage", handleSendMessage)
	http.HandleFunc("/getChats", handleGetChats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
