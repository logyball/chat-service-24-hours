package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetChatReturnsCorrectChat(t *testing.T) {
	type test struct {
		users   UserDb
		userOne uuid.UUID
		userTwo uuid.UUID
		result  *Chat
	}

	userIdOne := uuid.New()
	userIdTwo := uuid.New()

	chatOne := Chat{
		UserOneId: userIdOne,
		UserTwoId: userIdTwo,
		Messages:  []Message{},
	}

	emptyDatabase := UserDb{
		Users:         map[uuid.UUID]bool{},
		UsersToUsers:  map[uuid.UUID]uuid.UUID{},
		UserOneToChat: map[uuid.UUID]*Chat{},
	}

	fullDatabase := UserDb{
		Users: map[uuid.UUID]bool{
			userIdOne: true,
			userIdTwo: true,
		},
		UsersToUsers:  map[uuid.UUID]uuid.UUID{userIdOne: userIdTwo},
		UserOneToChat: map[uuid.UUID]*Chat{userIdOne: &chatOne},
	}

	tests := []test{
		{users: emptyDatabase, userOne: userIdOne, userTwo: userIdTwo, result: nil},
		{users: fullDatabase, userOne: userIdOne, userTwo: userIdTwo, result: &chatOne},
	}

	for _, test := range tests {
		got := getChatBetweenUsers(userIdOne, userIdTwo)
		assert.Equal(t, got, test.result)
	}
}
