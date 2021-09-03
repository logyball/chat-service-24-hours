package main

import (
	"log"
	"sync"
)

type UserDb struct {
	lock        sync.Mutex
	curId       int32
	Users       map[int32]bool
	UserToChats map[int32][]*Chat
}

func (u *UserDb) getUser(usr int32) bool {
	return u.Users[usr]
}

func (u *UserDb) addUser() int32 {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.curId++
	u.Users[u.curId] = true
	return u.curId
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
func (u *UserDb) getChatBetweenUsers(to int32, from int32) *Chat {
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
