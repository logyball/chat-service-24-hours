package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

type Message struct {
	Msg       string
	To        User
	From      User
	timestamp time.Time
}

type Chat struct {
	UserOne  User
	UserTwo  User
	Messages []Message
}

func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getTime() time.Time {
	return time.Now().UTC()
}

var userDatabase UserDb

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newUser" {
		http.Error(w, "Bad URL", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("%v Method not supported", r.Method), http.StatusNotFound)
		return
	}

	newUser := CreateNewUserAddToDb(&userDatabase)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func handleEmptyGetReq(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Bad URL", http.StatusNotFound)
		return
	}
	log.Printf("Got a GET request to the root at %v", getTime())
	fmt.Fprintf(w, "Hello\n")
}

func init() {
	userDatabase = UserDb{
		Users: make(map[uuid.UUID]User),
	}
}

func main() {
	u := User{
		Id: uuid.New(),
	}
	fmt.Printf("%v", u)

	http.HandleFunc("/", handleEmptyGetReq)
	http.HandleFunc("/newUser", handleNewUser)
	must(http.ListenAndServe(":8080", nil))
}
