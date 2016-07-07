# minimal linux distribution
FROM golang:1.6-wheezy

# GO and PATH env variables already set in golang image

# set the go path to import the source project
WORKDIR $GOPATH/src/github.com/sebastienfr/handsongo
ADD . $GOPATH/src/github.com/sebastienfr/handsongo

# In one command-line (for reduce memory usage purposes),
# we install the required software,
# we build handsongo program
# we clean the system from all build dependencies
RUN make all && rm -rf $GOPATH/pkg && rm -rf $GOPATH/src

# by default, the exposed ports are 8020 (HTTP)
EXPOSE 8020
