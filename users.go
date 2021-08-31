package main

import (
	"log"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	// other fields as necessary
}

type UserDb struct {
	UsersToChats map[uuid.UUID]map[uuid.UUID]*Chat
	/*
		{
			"user uuid" -> {
				"userTwo uuid" -> chat,
				"userThree uuid" -> chat
			}
		}
	*/
}

func CreateNewUserAddToDb(d *UserDb) User {
	u := uuid.New()
	newUser := User{
		Id: u,
	}
	emptyChats := make(map[uuid.UUID]*Chat)

	userDatabase.UsersToChats[u] = emptyChats
	log.Printf("Created a new user %v with no chats at %v", u, getTime())
	log.Printf("Users: %v", userDatabase)
	return newUser
}
