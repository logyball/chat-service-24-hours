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
	if m.From == uuid.Nil {
		return errors.New("no \"From\" User")
	}
	if !userDatabase.getUser(m.To) {
		return errors.New("Chat recipient doesnt exist")
	}
	if !userDatabase.getUser(m.From) {
		return errors.New("Chat sender doesnt exist")
	}
	return nil
}

func findExistingOrCreateNewChat(m *Message) *Chat {
	chat := userDatabase.getChatBetweenUsers(m.To, m.From)
	if chat != nil {
		log.Printf("chat found")
		chat.Messages = append(chat.Messages, *m)
		return chat
	}
	newChat := userDatabase.createChat(m)
	return newChat
}
