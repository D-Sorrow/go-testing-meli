all: test
test:
	go test ./... -coverprofile=coverage.out

template: test
	go tool cover -html=coverage.out -o coverage.html

cover-full:
	go tool cover -func=coverage.out