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

(With server running in a separate terminal)

Client:
```bash
$ ./test_setup/populate_db.sh 
{"Id":25}
{"Id":26}
{"Id":27}
{"Id":28}
{"Id":29}
{"Id":30}
{"Id":31}
{"Id":32}
{"Id":33}
{"Id":34}
{"Id":35}
{"Id":36}


Populating Chats...


Chats between: 4 and 3
["[2021-09-07 13:25:47.854411 +0000 UTC]\n- To: 4\n- From: 3\nMessage: Whats for Dinner","[2021-09-07 13:25:47.87887 +0000 UTC]\n- To: 4\n- From: 3\nMessage: Whats for Dinner","[2021-09-07 13:25:59.668891 +0000 UTC]\n- To: 3\n- From: 4\nMessage: Hello","[2021-09-07 13:25:59.769629 +0000 UTC]\n- To: 3\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:26:00.016134 +0000 UTC]\n- To: 3\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:26:00.248154 +0000 UTC]\n- To: 3\n- From: 4\nMessage: Are you there?","[2021-09-07 13:26:00.347486 +0000 UTC]\n- To: 4\n- From: 3\nMessage: Hello"]
Chats between: 2 and 8
["[2021-09-07 13:25:59.749417 +0000 UTC]\n- To: 8\n- From: 2\nMessage: Hello","[2021-09-07 13:25:59.905624 +0000 UTC]\n- To: 8\n- From: 2\nMessage: Goodbye","[2021-09-07 13:25:59.925875 +0000 UTC]\n- To: 8\n- From: 2\nMessage: Goodbye","[2021-09-07 13:26:00.407735 +0000 UTC]\n- To: 8\n- From: 2\nMessage: Are you there?","[2021-09-07 13:27:19.31761 +0000 UTC]\n- To: 2\n- From: 8\nMessage: ASL","[2021-09-07 13:27:19.374994 +0000 UTC]\n- To: 8\n- From: 2\nMessage: Hello","[2021-09-07 13:27:19.775485 +0000 UTC]\n- To: 2\n- From: 8\nMessage: Goodbye"]
Chats between: 6 and 2
["[2021-09-07 13:25:17.221713 +0000 UTC]\n- To: 2\n- From: 6\nMessage: Goodbye","[2021-09-07 13:27:19.466059 +0000 UTC]\n- To: 6\n- From: 2\nMessage: Goodbye","[2021-09-07 13:27:19.576232 +0000 UTC]\n- To: 6\n- From: 2\nMessage: ASL","[2021-09-07 13:27:19.684797 +0000 UTC]\n- To: 2\n- From: 6\nMessage: Goodbye"]
Chats between: 4 and 9
["[2021-09-07 13:27:19.529928 +0000 UTC]\n- To: 4\n- From: 9\nMessage: Goodbye","[2021-09-07 13:27:19.797236 +0000 UTC]\n- To: 9\n- From: 4\nMessage: u up?"]
Chats between: 7 and 7
null
Chats between: 5 and 8
["[2021-09-07 13:25:17.164226 +0000 UTC]\n- To: 8\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:25:59.713721 +0000 UTC]\n- To: 5\n- From: 8\nMessage: u up?","[2021-09-07 13:26:00.006358 +0000 UTC]\n- To: 5\n- From: 8\nMessage: Goodbye","[2021-09-07 13:27:19.365822 +0000 UTC]\n- To: 5\n- From: 8\nMessage: ASL","[2021-09-07 13:27:19.566852 +0000 UTC]\n- To: 8\n- From: 5\nMessage: Hello","[2021-09-07 13:27:19.728916 +0000 UTC]\n- To: 8\n- From: 5\nMessage: Goodbye","[2021-09-07 13:27:19.898739 +0000 UTC]\n- To: 8\n- From: 5\nMessage: Whats for Dinner"]
Chats between: 7 and 1
["[2021-09-07 13:25:59.659221 +0000 UTC]\n- To: 7\n- From: 1\nMessage: Whats for Dinner","[2021-09-07 13:25:59.825241 +0000 UTC]\n- To: 7\n- From: 1\nMessage: Hello","[2021-09-07 13:27:19.484249 +0000 UTC]\n- To: 7\n- From: 1\nMessage: Whats for Dinner","[2021-09-07 13:27:19.585539 +0000 UTC]\n- To: 7\n- From: 1\nMessage: Goodbye","[2021-09-07 13:27:20.176943 +0000 UTC]\n- To: 7\n- From: 1\nMessage: Are you there?","[2021-09-07 13:27:20.186208 +0000 UTC]\n- To: 7\n- From: 1\nMessage: u up?","[2021-09-07 13:27:20.2183 +0000 UTC]\n- To: 1\n- From: 7\nMessage: u up?"]
Chats between: 10 and 2
["[2021-09-07 13:25:59.416809 +0000 UTC]\n- To: 2\n- From: 10\nMessage: u up?","[2021-09-07 13:25:59.995528 +0000 UTC]\n- To: 10\n- From: 2\nMessage: Whats for Dinner","[2021-09-07 13:26:00.458369 +0000 UTC]\n- To: 10\n- From: 2\nMessage: ASL","[2021-09-07 13:27:19.539182 +0000 UTC]\n- To: 2\n- From: 10\nMessage: u up?","[2021-09-07 13:27:19.629496 +0000 UTC]\n- To: 2\n- From: 10\nMessage: Goodbye","[2021-09-07 13:27:19.806295 +0000 UTC]\n- To: 2\n- From: 10\nMessage: ASL","[2021-09-07 13:27:20.134553 +0000 UTC]\n- To: 2\n- From: 10\nMessage: Goodbye"]
Chats between: 7 and 5
["[2021-09-07 13:25:47.865734 +0000 UTC]\n- To: 5\n- From: 7\nMessage: u up?","[2021-09-07 13:25:59.587026 +0000 UTC]\n- To: 5\n- From: 7\nMessage: Goodbye","[2021-09-07 13:27:19.327045 +0000 UTC]\n- To: 7\n- From: 5\nMessage: u up?","[2021-09-07 13:27:19.47509 +0000 UTC]\n- To: 7\n- From: 5\nMessage: Goodbye","[2021-09-07 13:27:19.618676 +0000 UTC]\n- To: 5\n- From: 7\nMessage: Goodbye","[2021-09-07 13:27:19.719236 +0000 UTC]\n- To: 7\n- From: 5\nMessage: Goodbye","[2021-09-07 13:27:20.016132 +0000 UTC]\n- To: 7\n- From: 5\nMessage: Goodbye","[2021-09-07 13:27:20.027439 +0000 UTC]\n- To: 5\n- From: 7\nMessage: Are you there?"]
Chats between: 2 and 6
["[2021-09-07 13:25:17.221713 +0000 UTC]\n- To: 2\n- From: 6\nMessage: Goodbye","[2021-09-07 13:27:19.466059 +0000 UTC]\n- To: 6\n- From: 2\nMessage: Goodbye","[2021-09-07 13:27:19.576232 +0000 UTC]\n- To: 6\n- From: 2\nMessage: ASL","[2021-09-07 13:27:19.684797 +0000 UTC]\n- To: 2\n- From: 6\nMessage: Goodbye"]
```