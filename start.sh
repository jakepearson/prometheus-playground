#!/bin/bash

rm -rf ./data
alertmanager &
watcher -cmd "http post localhost:9093/-/reload" -dotfiles=false -list -recursive=false -keepalive &
prometheus --web.enable-lifecycle &
watcher -cmd "http post localhost:9090/-/reload" -dotfiles=false -list -recursive=false -keepalive &
gin run main.go