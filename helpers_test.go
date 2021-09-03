package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetChatReturnsCorrectChat(t *testing.T) {
	type test struct {
		users       *UserDb
		chatIsEmpty bool
	}

	emptyDatabase := UserDb{
		lock:        sync.Mutex{},
		Users:       map[int32]bool{},
		UserToChats: map[int32][]*Chat{},
	}

	fullDatabase := UserDb{
		lock:        sync.Mutex{},
		Users:       map[int32]bool{},
		UserToChats: map[int32][]*Chat{},
	}

	userIdOne := fullDatabase.addUser()
	userIdTwo := fullDatabase.addUser()
	msgOne := Message{
		Msg:       "asdf",
		To:        userIdOne,
		From:      userIdTwo,
		timestamp: time.Time{},
	}
	fullDatabase.createChat(&msgOne)

	tests := []test{
		{users: &emptyDatabase, chatIsEmpty: true},
		{users: &fullDatabase, chatIsEmpty: false},
	}

	for _, test := range tests {
		got := test.users.getChatBetweenUsers(userIdOne, userIdTwo)
		if test.chatIsEmpty {
			assert.Nil(t, got)
		} else {
			assert.NotNil(t, got)
		}
	}
}
