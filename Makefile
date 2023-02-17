default: run

run:
	go run statd.go

dist:
	rm -rf statd
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o statd -ldflags "-s -w" -trimpath statd.go