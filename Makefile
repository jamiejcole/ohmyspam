.PHONY: test server

test:
	python3 ./dev/send.py

server:
	cd backend && go run ./cmd/server/main.go
