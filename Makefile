LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

update-packages:
	go get -u ./...

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out