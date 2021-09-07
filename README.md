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


Populating Chats...




Chats between: 6 and 5


["[2021-09-07 13:33:47.210068 +0000 UTC]\n- To: 6\n- From: 5\nMessage: u up?","[2021-09-07 13:33:47.707948 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:48.953113 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Hello","[2021-09-07 13:33:49.183175 +0000 UTC]\n- To: 5\n- From: 6\nMessage: Are you there?","[2021-09-07 13:33:49.896808 +0000 UTC]\n- To: 6\n- From: 5\nMessage: ASL","[2021-09-07 13:33:50.463506 +0000 UTC]\n- To: 5\n- From: 6\nMessage: u up?","[2021-09-07 13:33:50.473332 +0000 UTC]\n- To: 6\n- From: 5\nMessage: ASL","[2021-09-07 13:33:51.5421 +0000 UTC]\n- To: 6\n- From: 5\nMessage: u up?","[2021-09-07 13:33:51.860836 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Are you there?","[2021-09-07 13:33:52.956941 +0000 UTC]\n- To: 5\n- From: 6\nMessage: Hello","[2021-09-07 13:33:53.751256 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Goodbye","[2021-09-07 13:33:53.961543 +0000 UTC]\n- To: 5\n- From: 6\nMessage: u up?","[2021-09-07 13:33:55.284191 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:55.58925 +0000 UTC]\n- To: 6\n- From: 5\nMessage: u up?","[2021-09-07 13:33:56.19091 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:56.611937 +0000 UTC]\n- To: 6\n- From: 5\nMessage: Hello","[2021-09-07 13:33:56.91001 +0000 UTC]\n- To: 5\n- From: 6\nMessage: Whats for Dinner","[2021-09-07 13:33:57.279928 +0000 UTC]\n- To: 5\n- From: 6\nMessage: Goodbye"]


Chats between: 1 and 9


["[2021-09-07 13:33:46.993574 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Goodbye","[2021-09-07 13:33:47.014515 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Whats for Dinner","[2021-09-07 13:33:47.089617 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Hello","[2021-09-07 13:33:47.188602 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Are you there?","[2021-09-07 13:33:47.623742 +0000 UTC]\n- To: 1\n- From: 9\nMessage: ASL","[2021-09-07 13:33:47.92854 +0000 UTC]\n- To: 9\n- From: 1\nMessage: ASL","[2021-09-07 13:33:48.308679 +0000 UTC]\n- To: 9\n- From: 1\nMessage: u up?","[2021-09-07 13:33:48.704254 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Hello","[2021-09-07 13:33:49.013244 +0000 UTC]\n- To: 9\n- From: 1\nMessage: u up?","[2021-09-07 13:33:49.727914 +0000 UTC]\n- To: 9\n- From: 1\nMessage: u up?","[2021-09-07 13:33:49.915253 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Hello","[2021-09-07 13:33:49.926012 +0000 UTC]\n- To: 1\n- From: 9\nMessage: ASL","[2021-09-07 13:33:50.35731 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Are you there?","[2021-09-07 13:33:51.415877 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Goodbye","[2021-09-07 13:33:51.745481 +0000 UTC]\n- To: 9\n- From: 1\nMessage: Hello","[2021-09-07 13:33:52.310253 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Are you there?","[2021-09-07 13:33:52.369545 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Whats for Dinner","[2021-09-07 13:33:52.51805 +0000 UTC]\n- To: 1\n- From: 9\nMessage: u up?","[2021-09-07 13:33:53.091862 +0000 UTC]\n- To: 9\n- From: 1\nMessage: Goodbye","[2021-09-07 13:33:53.444562 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Whats for Dinner","[2021-09-07 13:33:53.543638 +0000 UTC]\n- To: 9\n- From: 1\nMessage: ASL","[2021-09-07 13:33:54.463176 +0000 UTC]\n- To: 1\n- From: 9\nMessage: u up?","[2021-09-07 13:33:55.338573 +0000 UTC]\n- To: 9\n- From: 1\nMessage: ASL","[2021-09-07 13:33:55.55779 +0000 UTC]\n- To: 1\n- From: 9\nMessage: Hello","[2021-09-07 13:33:55.612457 +0000 UTC]\n- To: 9\n- From: 1\nMessage: Hello","[2021-09-07 13:33:56.244587 +0000 UTC]\n- To: 9\n- From: 1\nMessage: Hello","[2021-09-07 13:33:56.430598 +0000 UTC]\n- To: 9\n- From: 1\nMessage: Goodbye"]


Chats between: 5 and 4


["[2021-09-07 13:33:46.835094 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Are you there?","[2021-09-07 13:33:47.036074 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:47.671189 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Hello","[2021-09-07 13:33:48.013529 +0000 UTC]\n- To: 5\n- From: 4\nMessage: ASL","[2021-09-07 13:33:48.089336 +0000 UTC]\n- To: 5\n- From: 4\nMessage: u up?","[2021-09-07 13:33:48.28824 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Hello","[2021-09-07 13:33:49.611282 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:33:50.568941 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Are you there?","[2021-09-07 13:33:50.641511 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Hello","[2021-09-07 13:33:50.679106 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:33:50.867077 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:33:50.976052 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:51.209616 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:51.785691 +0000 UTC]\n- To: 5\n- From: 4\nMessage: u up?","[2021-09-07 13:33:52.196476 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:53.080348 +0000 UTC]\n- To: 4\n- From: 5\nMessage: u up?","[2021-09-07 13:33:54.152163 +0000 UTC]\n- To: 4\n- From: 5\nMessage: u up?","[2021-09-07 13:33:54.221018 +0000 UTC]\n- To: 4\n- From: 5\nMessage: u up?","[2021-09-07 13:33:55.410542 +0000 UTC]\n- To: 4\n- From: 5\nMessage: u up?","[2021-09-07 13:33:55.764625 +0000 UTC]\n- To: 4\n- From: 5\nMessage: Whats for Dinner","[2021-09-07 13:33:55.86255 +0000 UTC]\n- To: 5\n- From: 4\nMessage: ASL","[2021-09-07 13:33:56.060991 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Hello","[2021-09-07 13:33:56.752525 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Goodbye","[2021-09-07 13:33:57.078356 +0000 UTC]\n- To: 5\n- From: 4\nMessage: Whats for Dinner"]


Chats between: 6 and 10


["[2021-09-07 13:33:48.2011 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Hello","[2021-09-07 13:33:48.427355 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Are you there?","[2021-09-07 13:33:48.738739 +0000 UTC]\n- To: 6\n- From: 10\nMessage: u up?","[2021-09-07 13:33:49.641376 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Whats for Dinner","[2021-09-07 13:33:49.949635 +0000 UTC]\n- To: 6\n- From: 10\nMessage: ASL","[2021-09-07 13:33:51.151411 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Goodbye","[2021-09-07 13:33:51.765225 +0000 UTC]\n- To: 6\n- From: 10\nMessage: ASL","[2021-09-07 13:33:52.144649 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Hello","[2021-09-07 13:33:52.37949 +0000 UTC]\n- To: 6\n- From: 10\nMessage: Are you there?","[2021-09-07 13:33:52.418531 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Hello","[2021-09-07 13:33:52.467054 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Whats for Dinner","[2021-09-07 13:33:52.921876 +0000 UTC]\n- To: 6\n- From: 10\nMessage: ASL","[2021-09-07 13:33:52.985867 +0000 UTC]\n- To: 6\n- From: 10\nMessage: Whats for Dinner","[2021-09-07 13:33:53.12235 +0000 UTC]\n- To: 10\n- From: 6\nMessage: Goodbye","[2021-09-07 13:33:53.929585 +0000 UTC]\n- To: 10\n- From: 6\nMessage: ASL","[2021-09-07 13:33:54.230785 +0000 UTC]\n- To: 6\n- From: 10\nMessage: u up?","[2021-09-07 13:33:54.383279 +0000 UTC]\n- To: 6\n- From: 10\nMessage: Whats for Dinner","[2021-09-07 13:33:54.854427 +0000 UTC]\n- To: 6\n- From: 10\nMessage: ASL","[2021-09-07 13:33:55.016196 +0000 UTC]\n- To: 6\n- From: 10\nMessage: Are you there?","[2021-09-07 13:33:57.216655 +0000 UTC]\n- To: 6\n- From: 10\nMessage: Goodbye"]


Chats between: 7 and 4


["[2021-09-07 13:33:48.046979 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Goodbye","[2021-09-07 13:33:48.133226 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Goodbye","[2021-09-07 13:33:48.493887 +0000 UTC]\n- To: 4\n- From: 7\nMessage: Whats for Dinner","[2021-09-07 13:33:48.922488 +0000 UTC]\n- To: 4\n- From: 7\nMessage: ASL","[2021-09-07 13:33:49.837725 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Hello","[2021-09-07 13:33:50.08959 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Hello","[2021-09-07 13:33:50.888172 +0000 UTC]\n- To: 4\n- From: 7\nMessage: ASL","[2021-09-07 13:33:52.45838 +0000 UTC]\n- To: 4\n- From: 7\nMessage: Hello","[2021-09-07 13:33:52.718192 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Hello","[2021-09-07 13:33:52.879716 +0000 UTC]\n- To: 7\n- From: 4\nMessage: ASL","[2021-09-07 13:33:53.247287 +0000 UTC]\n- To: 4\n- From: 7\nMessage: Goodbye","[2021-09-07 13:33:54.242184 +0000 UTC]\n- To: 4\n- From: 7\nMessage: ASL","[2021-09-07 13:33:54.632763 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Hello","[2021-09-07 13:33:54.718607 +0000 UTC]\n- To: 4\n- From: 7\nMessage: u up?","[2021-09-07 13:33:54.798691 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Whats for Dinner","[2021-09-07 13:33:55.059669 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Goodbye","[2021-09-07 13:33:55.784788 +0000 UTC]\n- To: 4\n- From: 7\nMessage: Goodbye","[2021-09-07 13:33:56.35762 +0000 UTC]\n- To: 7\n- From: 4\nMessage: u up?","[2021-09-07 13:33:56.52654 +0000 UTC]\n- To: 7\n- From: 4\nMessage: Hello","[2021-09-07 13:33:56.71861 +0000 UTC]\n- To: 4\n- From: 7\nMessage: ASL"]
```