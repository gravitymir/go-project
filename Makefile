.PHONY: build
build:
		go build -o z_api -v ./
		./z_api

.PHONY: test
test:
#go test -v -race -timeout 30s ./...
		go test -v -race -timeout 30s ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

.DEFAULT_GOAL := build



# build:
# 	docker-compose build todo-app

# run:
# 	docker-compose up todo-app

# test:
# 	go test -v ./...

# migrate:
# 	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

# swag:
# 	swag init -g cmd/main.go