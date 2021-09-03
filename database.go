package main

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type UserDb struct {
	lock        sync.Mutex
	Users       map[uuid.UUID]bool
	UserToChats map[uuid.UUID][]*Chat
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

	u.UserToChats[initMsg.From] = append(u.UserToChats[initMsg.From], &newChat)
	u.UserToChats[initMsg.To] = append(u.UserToChats[initMsg.To], &newChat)

	log.Printf("new Chat created")
	u.lock.Unlock()
	return &newChat
}

// TODO - optimize... this is awful
func (u *UserDb) getChatBetweenUsers(to uuid.UUID, from uuid.UUID) *Chat {
	var shorterChatList []*Chat

	toChats := u.UserToChats[to]
	fromChats := u.UserToChats[from]

	if len(toChats) > len(fromChats) {
		shorterChatList = fromChats
	} else {
		shorterChatList = toChats
	}

	for _, c := range shorterChatList {
		if (c.InitialFromUser == from && c.InitialToUser == to) || (c.InitialFromUser == to && c.InitialToUser == from) {
			return c
		}
	}

	// no chat exists
	return nil
}
