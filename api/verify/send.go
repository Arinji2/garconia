package verify

import (
	"context"
	_ "embed"
	"fmt"

	email_handler "github.com/arinji2/garconia/emails/templates"
	"github.com/arinji2/garconia/routes"
	"github.com/arinji2/garconia/sqlite"
	"github.com/arinji2/garconia/utils"
)

func VerifyEmail(con *sqlite.Connection, userData *sqlite.UserData) error {
	id, err := routes.GenerateID(context.Background(), con)
	if err != nil {
		return err
	}
	query := "INSERT INTO verifications (id, user_id) VALUES (?, ?)"
	_, err = con.Exec(context.Background(), query, id, userData.ID)
	if err != nil {
		return err
	}

	fmt.Println("Sending Verifcation Email to ", userData.Email)
	data := email_handler.VerifyEmailData{
		VerificationURL: utils.GenerateVerificationURL(id),
	}
	err = email_handler.SendVerificationEmail(data, userData.Email)
	if err != nil {
		return err
	}

	return nil
}
