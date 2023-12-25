.PHONY: proto tidy vendor

tidy:
	go mod tidy

vendor:
	go mod vendor

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/application/*.proto

build:
	go build -o bin/server ./cmd/server 
	go build -o bin/client_load ./cmd/client_load 
