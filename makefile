build:
	mkdir -p functions
	go get ./...
	@go build -o functions/submit services/submit/main.go
	@go build -o functions/hello services/hello/main.go
