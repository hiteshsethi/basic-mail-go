package basicmail

import (
	"net/smtp"
	"fmt"
	"strings"
)

func SendEmail(toEmailStr string, fromEmail string, subject string, body string, host string,port int) error {

	c, err := smtp.Dial(addr(host, port))
	if err != nil {
		return err
	}
	if err := c.Mail(fromEmail); err != nil {
		return err
	}
	toEmail := strings.Split(toEmailStr, ",")

	for _, addr := range toEmail {
		if err := c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	tempSubject := "Subject: " + subject + "\n"
	msg := tempSubject + mime + body

	if _, err = fmt.Fprintf(w, msg); err != nil {
		w.Close()
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}

func addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
