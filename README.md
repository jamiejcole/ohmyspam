# ohmyspam
Selfhosted solution for spam/trash mail

## Getting Started

### Setup
1. Create a `.env` file based on `.env.example`
2. Install Go dependencies with `cd backend && go mod tidy`

### SQLite
- DB file lives in `data/db.sqlite`
- Override with `DB_PATH` in `.env`


### Running

**Start SMTP server:**
Either use `make server`, or manually with `cd backend && go run ./cmd/server/main.go`

**Send a test email:**
Send with `make test`

**Build executable:**
Use `make build` to create `bin/server`


### Env config
| Variable | Default | Values | Description |
---------------------------------------------
| MAIL_PORT | 2525 | xxxxx | Port on which the mail server runs |
| DB_PATH | ../data/db.sqlite | <path> | Path to store db at |
| AUTO_DELETE | true | `true|false` | Mail gets deleted automatically, or stored until manual deletion |
| DELETE_AFTER_VIEW | true | `true|false` | Determines whether mail deletion timer is started at view time, or at receiving time.  |
| MAIL_RETENTION | 5 | in minutes | Time after which mail is deleted |
| MAIL_DOMAIN | example.com | a domain | Domain at which to receive mail from