package emails

import (
	"os"

	"github.com/resend/resend-go/v2"
)

func SendEmail(email string, to string, title string) (string, error) {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	params := &resend.SendEmailRequest{
		From:    "Garconia <no-reply@mail.garconia.net>",
		To:      []string{to},
		Subject: title,
		Html:    email,
	}
	res, err := client.Emails.Send(params)
	if err != nil {
		return "", err
	}
	return res.Id, nil
}
