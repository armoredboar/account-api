package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"regexp"
)

const (
	mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// ValidateEmail checks if the specified email is in a valid format.
func ValidateEmail(email string) bool {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

// SendEmail to the specified address.
func SendEmail(address string, subject string, body string) {

	sender := os.Getenv("AB_EMAIL_NOREPLY")
	name := os.Getenv("AB_EMAIL_NOREPLY_NAME")
	password := os.Getenv("AB_EMAIL_NOREPLY_PASSWORD")
	host := os.Getenv("AB_EMAIL_SMTP")
	port := os.Getenv("AB_EMAIL_PORT")

	from := mail.Address{name, sender}
	to := mail.Address{"", address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += mime + "\r\n" + body

	auth := smtp.PlainAuth("", sender, password, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", host+":"+port, tlsconfig)
	if err != nil {
		log.Panic(err.Error() + " - " + host + ":" + port)
	}

	// Instantiates a new smtp client.
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Authenticate the account.
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// Set the from address.
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	// Set the to address.
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Create the buffer to write the message.
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	// Write the message to que buffer.
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	// Close the buffer.
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	// Close the client.
	c.Quit()
}
