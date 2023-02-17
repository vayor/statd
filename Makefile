default: run

run:
	go run statd.go

dist:
	rm -rf statd
	go build -o statd -ldflags "-s -w" -trimpath statd.go