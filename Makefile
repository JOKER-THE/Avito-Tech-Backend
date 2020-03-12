.PHONY: build
build:
	go get -u github.com/go-sql-driver/mysql
	go run index.go

.PHONY: start
start:
	go run index.go

.DEFAULT_GOAL := start