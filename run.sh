#!/bin/bash

CompileDaemon -build="go build server.go" -command="./server" -exclude-dir=".git" -exclude-dir="static" -exclude-dir="node_modules" -exclude-dir="bower_components"
