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

type ConfirmationEmailData struct {
	DeletionURL string
}

type VerifyEmailData struct {
	VerificationURL string
}

func SendVerificationEmail(data VerifyEmailData, to string) error {
	email := emails.EmailTemplateUtility(data, "verify", verifyEmailString)
	err := emails.SendEmail(email, to, "Verify Your Email For Garconia")
	if err != nil {
		return err
	}
	return nil
}

func SendConfirmationEmail(data ConfirmationEmailData, to string) error {
	email := emails.EmailTemplateUtility(data, "confirmation", confirmationEmailString)
	err := emails.SendEmail(email, to, "Confirmation Email for Garconia")
	if err != nil {
		return err
	}
	return nil
}

func SendReleaseEmail(to string) error {
	err := emails.SendEmail(releaseEmailString, to, "Garconia Released Today!!")
	if err != nil {
		return err
	}
	return nil
}
