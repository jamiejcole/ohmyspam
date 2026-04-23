package main

import (
	"fmt"
	"io"
	"log"

	"github.com/emersion/go-smtp"
)

type Backend struct{}

func (b *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	from string
	to   []string
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	_ = opts
	s.to = append(s.to, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	msg, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("==== NEW EMAIL ====")
	fmt.Println("From:", s.from)
	fmt.Println("To:", s.to)
	fmt.Println("Body:")
	fmt.Println(string(msg))
	fmt.Println("===================")

	return nil
}

func (s *Session) Reset()        {}
func (s *Session) Logout() error { return nil }

func main() {
	be := &Backend{}

	server := smtp.NewServer(be)

	server.Addr = ":2525"
	server.Domain = "localhost"
	server.AllowInsecureAuth = true

	log.Println("SMTP server running on :2525")
	log.Fatal(server.ListenAndServe())
}
