package main

import (
	"log"

	"github.com/google/uuid"
)

func CreateNewUserAddToDb(d *UserDb) User {
	u := uuid.New()
	newUser := User{
		Id: u,
	}

	userDatabase.Users[u] = true
	log.Printf("Created a new user %v with no chats at %v", u, getTime())
	log.Printf("Users: %v", userDatabase)
	return newUser
}
