#!/bin/bash

echo "Showing program first.go\n\n"
cat First/first.go
sleep 20
go run First/first.go
sleep 10
echo "Showing program second.go\n\n"
cat Second/second.go
sleep 20
go run Second/second.go
sleep 10
echo "Showing program third.go\n\n"
cat Third/third.go
sleep 20
go run Third/third.go
sleep 10
echo "Showing program fourth.go\n\n"
cat Fourth/fourth.go
sleep 20
go run Fourth/fourth.go
sleep 10

echo "All program outputs have been printed!"
