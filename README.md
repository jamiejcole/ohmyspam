# ohmyspam
Selfhosted solution for spam/trash mail

## Getting Started

### Setup
1. Create a `.env` file based on `.env.example`
2. Install Go dependencies with `cd backend && go mod tidy`

### Running

**Start SMTP server:**
Either use `make server`, or manually with `cd backend && go run ./cmd/server/main.go`

**Send a test email:**
Send with `make test`
