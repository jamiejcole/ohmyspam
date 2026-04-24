package mailserver

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/emersion/go-smtp"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Backend struct{}

func (b *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

func GenerateAddress() string {
	var colours = []string{"red", "black", "green", "yellow", "purple", "pink", "orange"}
	var animals = []string{"monkey", "cat", "dog", "wolf", "fish", "whale", "tiger"}

	return fmt.Sprintf("%s.%s",
		colours[rng.Intn(len(colours))],
		animals[rng.Intn(len(animals))],
	)
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
	fmt.Println(GenerateAddress())

	return nil
}

func (s *Session) Reset()        {}
func (s *Session) Logout() error { return nil }
