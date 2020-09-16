#!/bin/bash
kill -9 $(ps aux | grep "go" | awk '{print $2}')
git st && git br && git co -- . && git fetch origin && git pull origin master
nohup go run main.go >> request.log &
