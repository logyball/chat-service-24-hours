# chat-service-24-hours

A really stupid project.  Trying to hack together an HTTP-server-based chat service by myself in <24 hours. 

The results from <24 hours are saved off on the `24_hours` branch.  I can't reasonably have that code be the _first_ thing anyone sees while still touting myself as a "professional developer", so I've made some modifications to have this live in a less embarassing state.

[A full writeup of the experience is here](https://loganballard.com/index.php/2021/09/06/24-hours-in-go/).

# Compilation

```bash
$ make build
go test -v
=== RUN   TestGetChatReturnsCorrectChat
2021/09/07 08:55:14 new Chat created
--- PASS: TestGetChatReturnsCorrectChat (0.00s)
PASS
ok      example.com/chat_server 3.773s
go build -o ./bin/server *.go
```

This will result in a binary in `./bin/`.

# Running

`$ ./bin/server`

# Testing

With server running in a separate terminal:

Client:
```bash
`$ ./test_setup/populate_db.sh 
{"Id":1}
{"Id":2}
{"Id":3}
{"Id":4}
{"Id":5}
{"Id":6}
{"Id":7}
{"Id":8}
{"Id":9}
{"Id":10}
{"Id":11}
{"Id":12}
```

Server:
```bash
$ ./bin/server 
2021/09/07 09:21:44 Created a new user {1} with no chats at 2021-09-07 13:21:44.139238 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true]
2021/09/07 09:21:44 Created a new user {2} with no chats at 2021-09-07 13:21:44.148722 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true]
2021/09/07 09:21:44 Created a new user {3} with no chats at 2021-09-07 13:21:44.160341 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true]
2021/09/07 09:21:44 Created a new user {4} with no chats at 2021-09-07 13:21:44.171767 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true]
2021/09/07 09:21:44 Created a new user {5} with no chats at 2021-09-07 13:21:44.183283 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true]
2021/09/07 09:21:44 Created a new user {6} with no chats at 2021-09-07 13:21:44.19506 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true]
2021/09/07 09:21:44 Created a new user {7} with no chats at 2021-09-07 13:21:44.206385 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true]
2021/09/07 09:21:44 Created a new user {8} with no chats at 2021-09-07 13:21:44.217341 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true]
2021/09/07 09:21:44 Created a new user {9} with no chats at 2021-09-07 13:21:44.229695 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true 9:true]
2021/09/07 09:21:44 Created a new user {10} with no chats at 2021-09-07 13:21:44.240979 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true 9:true 10:true]
2021/09/07 09:21:44 Created a new user {11} with no chats at 2021-09-07 13:21:44.252135 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true 9:true 10:true 11:true]
2021/09/07 09:21:44 Created a new user {12} with no chats at 2021-09-07 13:21:44.263432 +0000 UTC
2021/09/07 09:21:44 Users: map[1:true 2:true 3:true 4:true 5:true 6:true 7:true 8:true 9:true 10:true 11:true 12:true]
2021/09/07 09:21:44 message: &{u up? 10 7 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{7 10 [{u up? 10 7 {275827000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Whats for Dinner 1 8 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{8 1 [{Whats for Dinner 1 8 {287612000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Hello 6 4 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{4 6 [{Hello 6 4 {297908000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Hello 4 1 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{1 4 [{Hello 4 1 {311052000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Whats for Dinner 4 3 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{3 4 [{Whats for Dinner 4 3 {323177000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Goodbye 4 3 {0 0 <nil>}}
2021/09/07 09:21:44 chat found
chat: &{3 4 [{Whats for Dinner 4 3 {323177000 63766617704 <nil>}} {Goodbye 4 3 {334729000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Hello 5 3 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{3 5 [{Hello 5 3 {347010000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{ASL 4 3 {0 0 <nil>}}
2021/09/07 09:21:44 chat found
chat: &{3 4 [{Whats for Dinner 4 3 {323177000 63766617704 <nil>}} {Goodbye 4 3 {334729000 63766617704 <nil>}} {ASL 4 3 {360156000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Are you there? 8 3 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{3 8 [{Are you there? 8 3 {371717000 63766617704 <nil>}}]}2021/09/07 09:21:44 message: &{Whats for Dinner 6 10 {0 0 <nil>}}
2021/09/07 09:21:44 new Chat created
chat: &{10 6 [{Whats for Dinner 6 10 {383148000 63766617704 <nil>}}]}
```