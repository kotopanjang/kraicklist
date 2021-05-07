cover-out: 
	go tool cover -func=./coverage/coverage.out

cover-html: 
	go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

run:
	go run ./main.go