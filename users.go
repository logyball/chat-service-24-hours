package main

import (
	"log"
)

func CreateNewUserAddToDb(d *UserDb) User {
	newUserId := d.addUser()
	newUser := User{
		Id: newUserId,
	}
	log.Printf("Created a new user %v with no chats at %v", newUser, getTime())
	log.Printf("Users: %v", d.Users)
	return newUser
}
