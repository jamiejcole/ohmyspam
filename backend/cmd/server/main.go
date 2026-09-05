package main

import (
	"fmt"
	"log"
	"log/slog"
	appdb "ohmyspam/internal/db"
	"ohmyspam/internal/mailserver"
	"os"
	"strconv"

	"github.com/emersion/go-smtp"
	"github.com/joho/godotenv"
)

func main() {
	// Setting up logging

	// TEMP
	removeTime := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		return a
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
		ReplaceAttr: removeTime,
	}))

	// Parsing env vars
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("warning: .env file not found, using default port 2525")
	}

	port := os.Getenv("MAIL_PORT")
	if port == "" {
		port = "2525"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "../data/db.sqlite"
	}

	// Checking and setting up db
	database, err := appdb.OpenAndInit(dbPath)
	if err != nil {
		logger.Error("Failed to initilise sqlite:", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	logger.Info("SQLite ready at:",
		slog.String("db path:", dbPath),
	)

	// Setting up mail server
	autoDelete, err := strconv.ParseBool(os.Getenv("AUTO_DELETE"))
	if err != nil {
		autoDelete = true
	}

	deleteAfterView, err := strconv.ParseBool(os.Getenv("DELETE_AFTER_VIEW"))
	if err != nil {
		deleteAfterView = true
	}

	mailRetentionPeriod := int32(5)
	if parsed, err := strconv.ParseInt(os.Getenv("MAIL_RETENTION"), 10, 32); err == nil {
		mailRetentionPeriod = int32(parsed)
	} else {
		logger.Info("Could not parse MAIL_RETENTION - defaulting to 5 minutes:", "error:", err)
	}

	mailDomain := os.Getenv("MAIL_DOMAIN")

	be := &mailserver.Backend{
		AutoDelete:          autoDelete,
		DeleteAfterView:     deleteAfterView,
		MailRetentionPeriod: mailRetentionPeriod,
	}
	s := smtp.NewServer(be)

	logger.Info("Server initialised with:",
		slog.String("AUTO_DELETE:", strconv.FormatBool(autoDelete)),
		slog.String("DELETE_AFTER_VIEW:", strconv.FormatBool(deleteAfterView)),
		slog.String("MAIL_RETENTION:", strconv.FormatInt(int64(mailRetentionPeriod), 10)),
	)

	logger.Info("Starting SMTP server on:",
		slog.String("Address", fmt.Sprintf(":%s", port)),
		slog.String("Domain", mailDomain),
	)

	s.Addr = fmt.Sprintf(":%s", port)
	s.Domain = mailDomain
	s.AllowInsecureAuth = true

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
