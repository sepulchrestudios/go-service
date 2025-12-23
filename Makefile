# Creates .env and .env.db-password files from their example files if they do not already exist
copy-env:
	cp -n .env.example .env
	cp -n .env.db-password.example .env.db-password

# Builds the protobuf files based on their specs
# https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc
proto:
	protoc -I ./proto \
		--go_out ./src/proto --go_opt paths=source_relative \
		--go-grpc_out ./src/proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./src/proto --grpc-gateway_opt paths=source_relative \
		./proto/*.proto

# Starts the Docker containers in detached mode after attempting to copy the env files
start: copy-env
	docker-compose up -d

# Stops and removes the Docker containers
stop:
	docker-compose down

.PHONY: copy-env proto start stop