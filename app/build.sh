#!/bin/bash

cp -R ./config ./build

go build -ldflags="-s -w" -o ./build/api.exe