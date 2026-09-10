package builder

//prod
import "fmt"

type Email struct {
	sender    string
	recipient string
	subject   string
	body      string
}

func (e Email) getSender() string {
	return e.sender
}

func (e Email) getRecipient() string {
	return e.recipient
}

func (e Email) getSubject() string {
	return e.subject
}

func (e Email) getBody() string {
	return e.body
}

func (e Email) String() string {
	return fmt.Sprintf("Email -> To: %s | From: %s | Subject: %s | Body: %s", e.recipient, e.sender, e.subject, e.body)
}
