package utilidades

import "net/smtp"

func SendEmail(email *NewEmail, message string) (bool, error) {
  // configure default data
	if len(email.HostSMTP) < 5 {
		email.HostSMTP = "smtp.gmail.com"
	}
	if len(email.PortSMTP) < 2 {
		email.PortSMTP = "587"
	}

  // create auth of mail
	var auth smtp.Auth = smtp.PlainAuth("", email.From, email.Password, email.HostSMTP)

  // send mail
	err := smtp.SendMail(email.HostSMTP+":"+email.PortSMTP, auth, email.From, email.To, []byte(message))
	if err != nil {
		return false, err
	}

	return true, err
}

// type of first param of the function
type NewEmail struct {
	From     string
	Password string
	To       []string
	Subject  string
	HostSMTP string
	PortSMTP string
}
