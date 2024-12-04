package verify

import (
	"context"
	_ "embed"
	"fmt"

	email_handler "github.com/arinji2/garconia/emails/templates"
	"github.com/arinji2/garconia/sqlite"
	"github.com/arinji2/garconia/utils"
)

func VerifyEmail(con *sqlite.Connection, userData *sqlite.UserData) error {
	id, err := utils.GenerateID(context.Background(), con, "signups")
	if err != nil {
		return err
	}

	fmt.Println("Sending Verifcation Email to ", userData.Email)
	data := email_handler.VerifyEmailData{
		VerificationURL: utils.GenerateVerificationURL(id),
	}
	email, err := email_handler.SendVerificationEmail(data, userData.Email)
	if err != nil {
		return err
	}

	query := "INSERT INTO verifications (id, user_id, resend_id) VALUES (?, ?, ?)"
	_, err = con.Exec(context.Background(), query, id, userData.ID, email)
	if err != nil {
		return err
	}

	return nil
}
