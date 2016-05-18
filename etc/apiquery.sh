#!/usr/bin/env bash
#debug
#set -x

# check if running on osx or linux
IP=127.0.0.1
if [ `uname -s` == "Darwin" ]; then
  IP=$(docker-machine ip)
fi

echo "Query all spirits..."
curl -s ${IP}:8020/spirits | jq

echo "Retrieve first spirit ID..."
ID=$(curl -s ${IP}:8020/spirits | jq '.[0]' | jq -r '.id')

echo "Query one spirit by found ID ${ID}"
curl -s ${IP}:8020/spirits/${ID} | jq
