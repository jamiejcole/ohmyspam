.PHONY: test server build

test:
	python3 ./dev/send.py

server:
	cd backend && go run ./cmd/server/main.go

build:
	mkdir -p ./bin
	cd backend && go build -o ../bin/server ./cmd/server
