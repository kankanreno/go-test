#!/bin/sh

APP="gomisc"

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main .
docker build -t $APP .
rm -f main

docker images
docker save $APP | gzip > $APP.tar.gz
docker rmi -f $APP

scp -C $APP.tar.gz root@v12:/root/$APP.tar.gz
rm -f $APP.tar.gz
