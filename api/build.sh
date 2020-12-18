!# /bin/bash
GOOS=linux go build -ldflags="-s -w" -o api ./api.go