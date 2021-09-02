package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	// other fields as necessary
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
	InitialFromUser uuid.UUID
	InitialToUser   uuid.UUID
	Messages        []Message
}
