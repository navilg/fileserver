#!/usr/bin/env bash

# Build
GOOS=linux go build -o ./bin/fileserver

# Copy to executable PATH
cp ./bin/fileserver $HOME/.local/bin/fileserver
chmod 750 $HOME/.local/bin/fileserver