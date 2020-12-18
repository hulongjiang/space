!# /bin/bash
GOOS=linux go build -ldflags="-s -w" -o userRpc ./user.go