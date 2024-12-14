package email_handler

import (
	_ "embed"

	"github.com/arinji2/garconia/emails"
)

//go:embed confirmation.html
var confirmationEmailString string

//go:embed release.html
var releaseEmailString string

//go:embed verify.html
var verifyEmailString string

type VerifyEmailData struct {
	VerificationURL string
}

func SendVerificationEmail(data VerifyEmailData, to string) (string, error) {
	email := emails.EmailTemplateUtility(data, "verify", verifyEmailString)
	res, err := emails.SendEmail(email, to, "Verify Your Email For Garconia")
	if err != nil {
		return "", err
	}
	return res, nil
}

func SendConfirmationEmail(to string) error {
	_, err := emails.SendEmail(confirmationEmailString, to, "Confirmation Email for Garconia")
	if err != nil {
		return err
	}
	return nil
}

func SendReleaseEmail(to string) error {
	_, err := emails.SendEmail(releaseEmailString, to, "Garconia Released Today!!")
	if err != nil {
		return err
	}
	return nil
}
