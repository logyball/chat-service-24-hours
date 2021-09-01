package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	// other fields as necessary
}

type UserDb struct {
	Users         map[uuid.UUID]bool
	UsersToUsers  map[uuid.UUID]uuid.UUID // user one -> user two, or "FROM" -> "TO"
	UserOneToChat map[uuid.UUID]*Chat     // user one = from, user two = to
}

type Message struct {
	Msg       string    `json:"Msg"`
	To        uuid.UUID `json:"To"`
	From      uuid.UUID `json:"From"`
	timestamp time.Time
}

type GetChat struct {
	To   uuid.UUID `json:"To"`
	From uuid.UUID `json:"From"`
}

type Chat struct {
	UserOneId uuid.UUID
	UserTwoId uuid.UUID
	Messages  []Message
}
