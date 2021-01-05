#!/bin/bash
LISTEN_ADDRESS=192.168.2.1:12345 STORAGE_ROOT=/tmp/1 go run ./dataServer.go&
LISTEN_ADDRESS=192.168.2.2:12345 STORAGE_ROOT=/tmp/2 go run ./dataServer.go&
LISTEN_ADDRESS=192.168.2.3:12345 STORAGE_ROOT=/tmp/3 go run ./dataServer.go&
LISTEN_ADDRESS=192.168.2.4:12345 STORAGE_ROOT=/tmp/4 go run ./dataServer.go&
LISTEN_ADDRESS=192.168.2.5:12345 STORAGE_ROOT=/tmp/5 go run ./dataServer.go&
LISTEN_ADDRESS=192.168.2.6:12345 STORAGE_ROOT=/tmp/6 go run ./dataServer.go&
