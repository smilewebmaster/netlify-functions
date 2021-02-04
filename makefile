##
## -- Makefile
##

SERVICES_INDEX = "./services/cmd/"

build: functions application

application:
	@echo "---> Application build start"
	@cd app && npm install && npm run build && npm run export
	@echo "---> Application build & export completed"

functions:
	@echo "---> Functions build start"
	@rm -rf bin/*
	@echo "---> Clear bin folder"
	@go get ./...
	# function list
	@go build -o bin/submit $(SERVICES_INDEX)/submit/main.go
	# end
	@echo "---> Functions Build Completed"

#
# -- END
#