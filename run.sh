#!/bin/bash

# build binary
go build || exit 1

# extract chat names
jq '.chats.list | keys[] as $k | "\(.[$k] | .name)"' --raw-output data/result.json > data/chats.txt

# generate all chats
while read chat; do
  echo "-> $chat"
  ./telegram-exporter -c "$chat" data/result.json > "data/$chat.csv" || rm "data/$chat.csv"
done <data/chats.txt