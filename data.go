package main

import (
	"time"
)

type User struct {
	Id int32
	// other fields as necessary
}

type Message struct {
	Msg       string `json:"Msg"`
	To        int32  `json:"To"`
	From      int32  `json:"From"`
	timestamp time.Time
}

type GetChat struct {
	To   int32 `json:"To"`
	From int32 `json:"From"`
}

type Chat struct {
	InitialFromUser int32
	InitialToUser   int32
	Messages        []Message
}
