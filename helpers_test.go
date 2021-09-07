package main

import (
	"io/ioutil"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setUp(disableOutput bool) {
	if disableOutput {
		log.SetOutput(ioutil.Discard)
	}
	for i := 0; i < 10; i++ {
		userDatabase.addUser()
	}
}

func clearDb() {
	userDatabase = UserDb{
		lock:        sync.Mutex{},
		curId:       0,
		Users:       map[int32]bool{},
		UserToChats: map[int32][]*Chat{},
	}
}

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

func TestValidateSendMessage(t *testing.T) {
	type test struct {
		input        Message
		errorPresent bool
	}

	tests := []test{
		{input: Message{Msg: ""}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			To:        1,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			From:      2,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			From:      2,
			To:        1,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "",
			To:        1,
			From:      2,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			To:        1,
			From:      1,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			To:        1,
			From:      100000,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			To:        1000000,
			From:      1,
			timestamp: time.Time{},
		}, errorPresent: true},
		{input: Message{
			Msg:       "asdf",
			To:        1,
			From:      2,
			timestamp: time.Time{},
		}, errorPresent: false},
	}

	setUp(false)
	defer clearDb()

	for _, te := range tests {
		got := validateSendMessage(&te.input)
		if te.errorPresent == (got == nil) {
			t.Fatalf("Expected to have error: %v, instead had: %v", te.errorPresent, got.Error())
		}
	}
}

func BenchmarkValidateSendMessage(b *testing.B) {
	input := Message{
		Msg:       "asdf",
		To:        1,
		From:      2,
		timestamp: time.Time{},
	}

	setUp(true)
	defer clearDb()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateSendMessage(&input)
	}
}

func BenchmarkValidateSendBadMessage(b *testing.B) {
	input := Message{
		Msg:       "asdf",
		To:        1,
		timestamp: time.Time{},
	}

	setUp(true)
	defer clearDb()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateSendMessage(&input)
	}
}
