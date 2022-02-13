#!/bin/sh

rm -rf bin

mkdir -p bin

GOOS=windows GOARCH=amd64 go build -o bin/commandlinequiz-amd64.exe

GOOS=windows GOARCH=386 go build -o bin/commandlinequiz-386.exe

GOOS=darwin GOARCH=amd64 go build -o bin/commandlinequiz-amd64-darwin

GOOS=linux GOARCH=amd64 go build -o bin/commandlinequiz-amd64-linux

GOOS=linux GOARCH=386 go build -o bin/commandlinequiz-386-linux