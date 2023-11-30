package email

import "gopkg.in/gomail.v2"

var (
	// dialer varieties
	sender string
	// global dialer
	dialer *gomail.Dialer
)

func Init(hostStr string, portStr int, name string, passStr string, nameStr string) {
	dialer = gomail.NewDialer(hostStr, portStr, name, passStr)
	sender = nameStr
}

func Send(dest []string, subject string, body string) error {
	message := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	message.SetHeaders(map[string][]string{
		"From":    {message.FormatAddress(dialer.Username, sender)},
		"To":      dest,
		"Subject": {subject},
	})
	message.SetBody("text/html", body)
	err := dialer.DialAndSend(message)
	return err
}
