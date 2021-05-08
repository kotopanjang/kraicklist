cover: 
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/coverage.out -covermode=atomic ./...
			go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
				go tool cover -func=./coverage/coverage.out

run:
	go run ./main.go