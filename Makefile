.PHONY: build

test-unit:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 sam build

