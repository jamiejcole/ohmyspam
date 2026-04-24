package main

import (
	"fmt"
	"log"
	"ohmyspam/internal/mailserver"
	"os"

	"github.com/emersion/go-smtp"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("warning: .env file not found, using default port 2525")
	}

	port := os.Getenv("MAIL_PORT")
	if port == "" {
		port = "2525"
	}

	be := &mailserver.Backend{}
	s := smtp.NewServer(be)

	s.Addr = fmt.Sprintf(":%s", port)
	s.Domain = "localhost"
	s.AllowInsecureAuth = true

	log.Printf("Starting SMTP server on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
