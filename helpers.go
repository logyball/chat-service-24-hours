package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func getTime() time.Time {
	return time.Now().UTC()
}

func getChatBetweenUsers(to uuid.UUID, from uuid.UUID) *Chat {
	checkUser := userDatabase.UsersToUsers[to]
	if checkUser != uuid.Nil {
		return userDatabase.UserOneToChat[to]
	}
	checkUser = userDatabase.UsersToUsers[from]
	if checkUser != uuid.Nil {
		return userDatabase.UserOneToChat[from]
	}
	// no chat exists
	return nil
}

func matchPathAndMethod(w http.ResponseWriter, r *http.Request, path string, method string) bool {
	if r.URL.Path != path {
		http.Error(w, "Bad URL", http.StatusNotFound)
		return false
	}
	if r.Method != method {
		http.Error(w, fmt.Sprintf("%v Method not supported", r.Method), http.StatusNotFound)
		return false
	}
	return true
}

func validateSendMessage(m *Message) error {
	log.Printf("message: %v", m)
	if m.Msg == "" {
		return errors.New("no message")
	}
	if m.To == uuid.Nil {
		return errors.New("no \"To\" User")
	}
	if !userDatabase.getUser(m.To) {
		return errors.New("Chat recipient doesnt exist")
	}
	if m.From == uuid.Nil {
		return errors.New("no \"From\" User")
	}
	return nil
}

func findExistingOrCreateNewChat(m *Message) *Chat {
	chat := getChatBetweenUsers(m.To, m.From)
	if chat != nil {
		log.Printf("chat found")
		chat.Messages = append(chat.Messages, *m)
		return chat
	}
	newChat := userDatabase.createChat(m)
	return newChat
}
