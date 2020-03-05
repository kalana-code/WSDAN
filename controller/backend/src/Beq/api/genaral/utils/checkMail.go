package utils

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

type smtpError struct {
	Err error
}

func (e smtpError) Error() string {
	return e.Err.Error()
}

func (e smtpError) Code() string {
	return e.Err.Error()[0:3]
}

func newsmtpError(err error) smtpError {
	return smtpError{
		Err: err,
	}
}

const forceDisconnectAfter = time.Second * 5

var (
	errBadFormat        = errors.New("invalid format")
	errUnresolvableHost = errors.New("unresolvable host")
	emailRegexp         = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func validateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return errBadFormat
	}
	return nil
}

func validateHost(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errUnresolvableHost
	}

	client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), forceDisconnectAfter)
	if err != nil {
		return newsmtpError(err)
	}
	defer client.Close()

	err = client.Hello("checkmail.me")
	if err != nil {
		return newsmtpError(err)
	}
	err = client.Mail("lansome-cowboy@gmail.com")
	if err != nil {
		return newsmtpError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return newsmtpError(err)
	}
	return nil
}

// DialTimeout returns a new Client connected to an SMTP server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func DialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}

func validate(email string) error {
	err1 := validateFormat(email)
	err2 := validateHost(email)
	if err1 != nil || err2 != nil {
		return errors.New("invalid format")
	}
	return nil
}
