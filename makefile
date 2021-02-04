build:
	mkdir -p functions
	go get ./...
	go build -o functions/xmachine-submit services/submit.go
