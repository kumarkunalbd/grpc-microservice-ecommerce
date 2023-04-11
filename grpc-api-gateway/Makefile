.PHONY: all
all:

.PHONY: proto-auth
proto-auth:
	@protoc pkg/auth/pb/*.proto --go_out=pkg/auth/pb --go_opt=paths=source_relative --go-grpc_out=pkg/auth/pb --go-grpc_opt=paths=source_relative --proto_path=pkg/auth/pb

.PHONY: proto-order
proto-order:
	@protoc pkg/order/pb/*.proto --go_out=pkg/order/pb --go_opt=paths=source_relative --go-grpc_out=pkg/order/pb --go-grpc_opt=paths=source_relative --proto_path=pkg/order/pb

.PHONY: proto-product
proto-product:
	@protoc pkg/product/pb/*.proto --go_out=pkg/product/pb --go_opt=paths=source_relative --go-grpc_out=pkg/product/pb --go-grpc_opt=paths=source_relative --proto_path=pkg/product/pb

.PHONY: server
server:
	@go run cmd/main.go