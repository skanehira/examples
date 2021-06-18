#!/bin/bash

cd server
go run main.go &

cd ../
go build -o out
./out
