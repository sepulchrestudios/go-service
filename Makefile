# Builds the protobuf files based on their specs
# https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc
proto:
	protoc -I ./proto \
		--go_out ./src/proto --go_opt paths=source_relative \
		--go-grpc_out ./src/proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./src/proto --grpc-gateway_opt paths=source_relative \
		./proto/*.proto

.PHONY: proto