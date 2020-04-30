all:
	go test ./... -v

format: 
	gofmt -s -w .
