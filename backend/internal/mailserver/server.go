package mailserver

import (
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/google/uuid"

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

	WriteToDB(s, string(msg))

	return nil
}

func WriteToDB(s *Session, msg string) error {
	db, err := sql.Open("sqlite", "../data/db.sqlite")

	if err != nil {
		fmt.Println(err)
		return err
	}

	sql := "INSERT INTO messages(id, mailbox_id, from_address, body_text) VALUES(?, ?, ?, ?)"

	_, err = db.Exec(sql, uuid.NewString(), s.to[0], s.from, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Session) Reset()        {}
func (s *Session) Logout() error { return nil }
