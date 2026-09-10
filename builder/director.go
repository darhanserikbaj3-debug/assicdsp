package builder

type EmailDirector struct{}

func NewEMailDirector() *EmailDirector {
	return &EmailDirector{}
}

func (d *EmailDirector) MakeWelcomeEmail(b EmailBuilder) {
	b.SetSender("darikPR@gmail.com").
		SetRecipient("newuser@gmail.com").
		SetSubject("Welcome here").
		SetBody("Hello there. I am exited ro have you here")
}

func (d *EmailDirector) MakePasswordResetEmail(b EmailBuilder) {
	b.SetSender("darikPR@gmail.com").
		SetRecipient("user@gmail.com").
		SetSubject("Reset Your Password").
		SetBody("Please click the link below to reset your password. The link expires in 15 minutes.")
}
