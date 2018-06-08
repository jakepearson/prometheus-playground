#!/bin/bash

rm -rf ./data
prometheus --web.enable-lifecycle &
watcher -cmd "http post localhost:9090/-/reload" -dotfiles=false -list -recursive=false -keepalive &
gin run main.go