package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Msg       string    `json:"Msg"`
	To        uuid.UUID `json:"To"`
	From      uuid.UUID `json:"From"`
	timestamp time.Time
}

type Chat struct {
	UserOneId uuid.UUID
	UserTwoId uuid.UUID
	Messages  []Message
}

func getTime() time.Time {
	return time.Now().UTC()
}

var userDatabase UserDb

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
	if _, ok := userDatabase.UsersToChats[m.To]; !ok {
		return errors.New("Chat recipient doesnt exist")
	}
	if m.From == uuid.Nil {
		return errors.New("no \"From\" User")
	}
	return nil
}

func findExistingOrCreateNewChat(m *Message) *Chat {
	if chat, found := userDatabase.UsersToChats[m.From][m.To]; found {
		log.Printf("chat found")
		chat.Messages = append(chat.Messages, *m)
		return chat
	}
	msgCopy := m
	newChat := Chat{
		UserOneId: m.From,
		UserTwoId: m.To,
		Messages:  []Message{*msgCopy},
	}
	toMap := map[uuid.UUID]*Chat{
		m.To: &newChat,
	}
	fromMap := map[uuid.UUID]*Chat{
		m.From: &newChat,
	}
	userDatabase.UsersToChats[m.From] = toMap
	userDatabase.UsersToChats[m.To] = fromMap
	log.Printf("new Chat created")
	return &newChat
}

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
	var m Message

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, fmt.Sprintf("Required json fields missing, or smthn lol: %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = validateSendMessage(&m)
	if err != nil {
		http.Error(w, fmt.Sprintf("Required json fields missing, or smthn lol: %v", err.Error()), http.StatusBadRequest)
		return
	}

	m.timestamp = getTime()

	c := findExistingOrCreateNewChat(&m)

	fmt.Printf("chat: %v", c)
}

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	if !matchPathAndMethod(w, r, "/newUser", "POST") {
		return
	}
	newUser := CreateNewUserAddToDb(&userDatabase)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func handleEmptyGetReq(w http.ResponseWriter, r *http.Request) {
	if !matchPathAndMethod(w, r, "/", "GET") {
		return
	}
	log.Printf("Got a GET request to the root at %v", getTime())
	fmt.Fprintf(w, "Hello\n")
}

func init() {

	userDatabase = UserDb{
		UsersToChats: make(map[uuid.UUID]map[uuid.UUID]*Chat),
	}
}

func main() {
	http.HandleFunc("/", handleEmptyGetReq)
	http.HandleFunc("/newUser", handleNewUser)
	http.HandleFunc("/sendMessage", handleSendMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
