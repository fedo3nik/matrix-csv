.PHONY: run test

run:
	go run .
test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
