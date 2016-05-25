#!/usr/bin/env bash
#debug
#set -x

# check if running on osx or linux
IP=127.0.0.1
if [ ! -z "$DOCKER_MACHINE_NAME" ]; then
  IP=$(docker-machine ip $DOCKER_MACHINE_NAME)
fi

echo "Create spirit from file..."
curl -s -X POST -H "Content-Type:application/json" -d @caroni.json ${IP}:8020/spirits | jq

echo "Query all spirits..."
curl -s ${IP}:8020/spirits | jq

echo "Retrieve first spirit by ID..."
ID=$(curl -s ${IP}:8020/spirits | jq '.[0]' | jq -r '.id')

echo "Query one spirit by found ID ${ID}"
curl -s ${IP}:8020/spirits/${ID} | jq

echo "Update spirit from file..."
curl -s -X PUT -H "Content-Type:application/json" -d @clairin.json ${IP}:8020/spirits/${ID} | jq

echo "Query one spirit by found ID ${ID} after update"
curl -s ${IP}:8020/spirits/${ID} | jq

echo "Deleting spirit by ID ${ID}"
curl -X DELETE -H "Content-Type:application/json" ${IP}:8020/spirits/${ID}
