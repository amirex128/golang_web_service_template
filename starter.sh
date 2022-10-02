#!/bin/bash
cd /home/sell/backend;
/usr/local/go/bin/go mod download;
/usr/local/go/bin/go mod vendor;
cd /home/sell/backend/cmd/server;
/usr/local/go/bin/go build main.go;
./main;

