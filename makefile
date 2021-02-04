##
## -- Makefile
##

SERVICES_INDEX = "./services/cmd/"

build:
	@rm -rf bin/*
	@go get ./...
	@go build -o bin/submit $(SERVICES_INDEX)/submit/main.go
	@go build -o bin/hello $(SERVICES_INDEX)/hello/main.go