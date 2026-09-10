package builder

import (
	"errors"
	"strings"
)

type ConcreteEmailBuilder struct {
	sender    string
	recipient string
	subject   string
	body      string
}

func NewConcreteEmailBuilder() *ConcreteEmailBuilder {
	return &ConcreteEmailBuilder{}
}

func (b *ConcreteEmailBuilder) SetSender(sender string) EmailBuilder {
	b.sender = sender
	return b
}

func (b *ConcreteEmailBuilder) SetRecipient(recipient string) EmailBuilder {
	b.recipient = recipient
	return b
}

func (b *ConcreteEmailBuilder) SetSubject(subject string) EmailBuilder {
	b.subject = subject
	return b
}

func (b *ConcreteEmailBuilder) SetBody(body string) EmailBuilder {
	b.body = body
	return b
}

func (b *ConcreteEmailBuilder) GetResult() (*Email, error) {
	if strings.TrimSpace(b.sender) == "" {
		return nil, errors.New("Validation error: no sender had found")
	}
	if strings.TrimSpace(b.recipient) == "" {
		return nil, errors.New("Validation error: no resipient had found")
	}
	if strings.TrimSpace(b.subject) == "" {
		return nil, errors.New("Validation error: no subject had found")
	}

	return &Email{
		sender:    b.sender,
		recipient: b.recipient,
		subject:   b.subject,
		body:      b.body,
	}, nil
}
