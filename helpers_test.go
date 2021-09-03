package main

import (
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetChatReturnsCorrectChat(t *testing.T) {
	type test struct {
		users       *UserDb
		chatIsEmpty bool
	}

	userIdOne := uuid.New()
	userIdTwo := uuid.New()

	msgOne := Message{
		Msg:       "asdf",
		To:        userIdTwo,
		From:      userIdOne,
		timestamp: time.Time{},
	}

	emptyDatabase := UserDb{
		lock:        sync.Mutex{},
		Users:       map[uuid.UUID]bool{},
		UserToChats: map[uuid.UUID][]*Chat{},
	}

	fullDatabase := UserDb{
		lock:        sync.Mutex{},
		Users:       map[uuid.UUID]bool{},
		UserToChats: map[uuid.UUID][]*Chat{},
	}

	fullDatabase.addUser(userIdOne)
	fullDatabase.addUser(userIdTwo)
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
