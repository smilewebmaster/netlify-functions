##
## -- Makefile
##

SERVICES_INDEX = "./services/cmd/"

build:
	@rm -rf bin/*
	@go get ./...
	@go build -o bin/submit $(SERVICES_INDEX)/submit/main.go
	@cd app
	@npm run build
	@npm run export