package builder

// inter
type EmailBuilder interface {
	SetSender(sender string) EmailBuilder
	SetRecipient(recipient string) EmailBuilder
	SetSubject(subject string) EmailBuilder
	SetBody(body string) EmailBuilder
}
