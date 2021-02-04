##
## -- Makefile
##

SERVICES_INDEX = "./services/cmd/"

build:
	@rm -rf bin/*
	@go get ./...
	@go build -o bin/submit $(SERVICES_INDEX)/submit/main.go
	@echo "---"
	@echo "Functions Build Completed"
	@echo "---"
	@SLEEP 5
	@cd app && npm install && npm run build && npm run export
	@echo "---"
	@echo "Application build & export completed"
	@echo "---"