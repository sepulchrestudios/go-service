# Creates .env and .env.db-password files from their example files if they do not already exist
copy-env:
	cp -n .env.example .env
	cp -n ./build/secrets/.env.db-password.example ./build/secrets/.env.db-password
	cp -n ./build/secrets/.env.feature-flag-sdk-key.example ./build/secrets/.env.feature-flag-sdk-key
	cp -n ./build/secrets/.env.mail-password.example ./build/secrets/.env.mail-password

# Builds the protobuf files based on their specs
# https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc
proto:
	protoc -I ./proto \
		--go_out ./src/proto --go_opt paths=source_relative \
		--go-grpc_out ./src/proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./src/proto --grpc-gateway_opt paths=source_relative \
		./proto/*.proto

# Builds the Go server Docker image without using cache
build: copy-env
	docker-compose build --no-cache go-server

# Starts the Docker containers in detached mode after attempting to copy the env files
start: copy-env
	docker-compose up -d

# Builds the Go server Docker image and starts the containers in one step for development
dev: build start

# Stops and removes the Docker containers
stop:
	docker-compose down

.PHONY: build copy-env dev proto start stop