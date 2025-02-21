#!/bin/sh

BINARY='/usr/local/bin'
APP=max_mtime

echo "Building $APP"
go build -ldflags="-s -w" $APP.go

echo "Installing $APP to $BINARY"
install $APP $BINARY

echo "Removing the build"
rm $APP