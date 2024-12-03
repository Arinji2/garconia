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

func VerificationEmail(data VerifyEmailData, to string) error {
	email := emails.EmailTemplateUtility(data, "verify", verifyEmailString)
	emails.SendEmail(email, to, "Verify Your Email")
	return nil
}

func ConfirmationEmail(data ConfirmationEmailData, to string) error {
	return nil
}

func ReleaseEmail(to string) error {
	return nil
}
