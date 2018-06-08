#!/bin/bash

brew install prometheus

echo "Installing Gin"
go get github.com/codegangsta/gin

echo "Installing Watcher"
go get -u github.com/radovskyb/watcher/...

echo "Done ... Happy Metrics-ing"