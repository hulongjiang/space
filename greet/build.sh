!# /bin/bash
GOOS=linux go build -ldflags="-s -w" -o greetApi ./greet.go