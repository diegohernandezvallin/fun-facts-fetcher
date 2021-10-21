#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o bin/main main.go
zip deploy/main.zip main