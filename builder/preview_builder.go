package builder

import (
	"errors"
	"fmt"
	"strings"
)

type EmailPreviewBuilder struct {
	sb strings.Builder
}

func NewEmailPreviewBuilder() *EmailPreviewBuilder {
	b := &EmailPreviewBuilder{}
	b.sb.WriteString("----------- EMAIL PREVIEW ---------\n")
	return b
}

func (b *EmailPreviewBuilder) SetSender(sender string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("FROM:     %s\n", sender))
	return b
}

func (b *EmailPreviewBuilder) SetRecipient(recipient string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("TO:      %s\n", recipient))
	return b
}

func (b *EmailPreviewBuilder) SetSubject(subject string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("SUBJECT: %s\n", subject))
	return b
}

func (b *EmailPreviewBuilder) SetBody(body string) EmailBuilder {
	b.sb.WriteString("-----------------------------------------------\n")
	b.sb.WriteString(fmt.Sprintf("%s\n", body))
	b.sb.WriteString("-----------------------------------------------")
	return b
}

func (b *EmailPreviewBuilder) GetResult() (string, error) {
	result := b.sb.String()
	if !strings.Contains(result, "TO:") || !strings.Contains(result, "FROM:") {
		return "", errors.New("validation error: preview missing headers")
	}
	return result, nil
}
