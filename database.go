package main

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type UserDb struct {
	lock                      sync.Mutex
	Users                     map[uuid.UUID]bool
	UsersToUsers              map[uuid.UUID]uuid.UUID // user one -> user two, or "FROM" -> "TO"
	MapInitialFromUserToChats map[uuid.UUID][]*Chat   // user one = from, user two = to
	MapInitialToUserToChats   map[uuid.UUID][]*Chat
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
		InitialFromUser: initMsg.From,
		InitialToUser:   initMsg.To,
		Messages:        []Message{*msgCopy},
	}
	u.MapInitialFromUserToChats[initMsg.From] = append(u.MapInitialFromUserToChats[initMsg.From], &newChat)
	u.MapInitialToUserToChats[initMsg.To] = append(u.MapInitialToUserToChats[initMsg.To], &newChat)
	u.UsersToUsers[initMsg.From] = initMsg.To
	u.UsersToUsers[initMsg.To] = initMsg.From
	log.Printf("new Chat created")
	u.lock.Unlock()
	return &newChat
}

// TODO - optimize... this is awful
func (u *UserDb) getChatBetweenUsers(to uuid.UUID, from uuid.UUID) *Chat {
	foundUserTo := u.UsersToUsers[from]
	if foundUserTo != uuid.Nil {
		for _, c := range u.MapInitialFromUserToChats[from] {
			if (c.InitialFromUser == from && c.InitialToUser == to) || (c.InitialFromUser == to && c.InitialToUser == from) {
				return c
			}
		}
		for _, c := range u.MapInitialToUserToChats[from] {
			if (c.InitialFromUser == from && c.InitialToUser == to) || (c.InitialFromUser == to && c.InitialToUser == from) {
				return c
			}
		}
	}
	// no chat exists
	return nil
}
