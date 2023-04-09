.PHONY: all
all:

.PHONY: proto-auth
proto-auth:
	@protoc pkg/auth/pb/*.proto --go_out=pkg/auth/pb --go_opt=paths=source_relative --go-grpc_out=pkg/auth/pb --go-grpc_opt=paths=source_relative --proto_path=pkg/auth/pb

.PHONY: server
server:
	@go run cmd/main.go