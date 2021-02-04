##
## -- Makefile
##

SERVICES_INDEX = "./services/cmd/"
UCYAN=\033[4;36m
NC=\033[0m

build: functions application

application:
	@echo "${UCYAN} ---> Application build start"
	@rm -rf app/.next
	@rm -rf app/out
	@cd app && npm install && npm run build && npm run export
	@echo "${UCYAN} ---> Application build & export completed"

functions:
	@echo "${UCYAN} ---> Functions build start"
	@rm -rf bin/*
	@echo "${UCYAN} ---> Clear bin folder"
	@go get ./...
	@go build -o bin/submit $(SERVICES_INDEX)/submit/main.go
	@echo "${UCYAN} ---> Functions Build Completed"

#
# -- END
#