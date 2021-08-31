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
	Users map[uuid.UUID]User
}

func CreateNewUserAddToDb(d *UserDb) User {
	u := uuid.New()
	newUser := User{
		Id: u,
	}

	userDatabase.Users[u] = newUser
	log.Printf("Created a new user %v at %v", u, getTime())
	log.Printf("Users: %v", userDatabase)
	return newUser
}
