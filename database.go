package main

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type UserDb struct {
	lock          sync.Mutex
	Users         map[uuid.UUID]bool
	UsersToUsers  map[uuid.UUID]uuid.UUID // user one -> user two, or "FROM" -> "TO"
	UserOneToChat map[uuid.UUID]*Chat     // user one = from, user two = to
}

func (u *UserDb) getUser(usr uuid.UUID) bool {
	return u.Users[usr]
}

func (u *UserDb) addUser(usr uuid.UUID) bool {
	if u.Users[usr] {
		log.Printf("User %v already exists", usr)
		return false
	}
	u.Users[usr] = true
	return true
}

func (u *UserDb) createChat(initMsg *Message) *Chat {
	u.lock.Lock()
	msgCopy := initMsg
	newChat := Chat{
		UserOneId: initMsg.From,
		UserTwoId: initMsg.To,
		Messages:  []Message{*msgCopy},
	}
	userDatabase.UserOneToChat[initMsg.From] = &newChat
	userDatabase.UsersToUsers[initMsg.From] = initMsg.To
	log.Printf("new Chat created")
	u.lock.Unlock()
	return &newChat
}
