# minimal linux distribution
FROM golang:1.6-alpine

# GO and PATH env variables already set in golang image
# to reduce download time
RUN apk add -U make

# set the go path to import the source project
WORKDIR $GOPATH/src/github.com/sebastienfr/handsongo
ADD . $GOPATH/src/github.com/sebastienfr/handsongo

# In one command-line (for reduce memory usage purposes),
# we install the required software,
# we build handsongo program
# we clean the system from all build dependencies
RUN make all && apk del make && \
  rm -rf /gopath/pkg && \
  rm -rf /gopath/src && \
  rm -rf /var/cache/apk/*

# by default, the exposed ports are 8020 (HTTP)
EXPOSE 8020
