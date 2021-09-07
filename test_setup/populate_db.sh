#!/bin/bash

# A stupid script that adds some stuff to the DB so that there's some data to use

### Create users 1-12
for _ in {1..12}; do
    curl -X POST localhost:8080/newUser
done

# Populate Chats

printf "\n\nPopulating Chats...\n\n\n"

RAND_MESSAGES=("Hello" "Whats for Dinner" "Are you there?" "ASL" "u up?" "Goodbye")
for _ in {1..100}; do
    USER_1=$(($RANDOM % 10 + 1))
    USER_2=$(($RANDOM % 10 + 1))
    while [ "$USER_1" = "$USER_2" ] ; do
        USER_2=$(($RANDOM % 10 + 1))
    done
    MSG=${RAND_MESSAGES[$(($RANDOM % 6))]}
    
    curl -X POST \
        -H "content-type: Application/Json" \
        -d "{\"To\": $USER_1, \"From\": $USER_2, \"Msg\": \"$MSG\"}" \
        localhost:8080/sendMessage

done

for _ in {1..10}; do
    USER_1=$(($RANDOM % 10 + 1))
    USER_2=$(($RANDOM % 10 + 1))

    echo "Chats between: $USER_1 and $USER_2"

    curl -X GET \
        -H "content-type: Application/Json" \
        -d "{\"To\": $USER_1, \"From\": $USER_2}" \
        localhost:8080/getChats
done