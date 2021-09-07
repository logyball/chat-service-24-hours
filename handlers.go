package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
	var m Message

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = validateSendMessage(&m)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err.Error()), http.StatusBadRequest)
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

func handleGetChats(w http.ResponseWriter, r *http.Request) {
	if !matchPathAndMethod(w, r, "/getChats", "GET") {
		return
	}
	var getChatMsg GetChat

	err := json.NewDecoder(r.Body).Decode(&getChatMsg)

	if err != nil {
		http.Error(w, fmt.Sprintf("Required json fields missing, or smthn lol: %v", err.Error()), http.StatusBadRequest)
		return
	}

	chatBetweenUsers := userDatabase.getChatBetweenUsers(getChatMsg.To, getChatMsg.From)
	var messages []string
	if chatBetweenUsers != nil {
		for _, m := range chatBetweenUsers.Messages {
			messages = append(messages, fmt.Sprintf("[%v]\n- To: %v\n- From: %v\nMessage: %v", m.timestamp, m.To, m.From, m.Msg))
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}
