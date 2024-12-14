package routes

import (
	"context"

	"github.com/arinji2/garconia/sqlite"
	"github.com/arinji2/garconia/utils"
)

type AddEmailRequest struct {
	Email string `json:"email"`
}

func addEmail(ctx context.Context, email string, con *sqlite.Connection) (string, error) {
	id, err := utils.GenerateID(ctx, con, "signups")
	if err != nil {
		return "", err
	}
	query := "INSERT INTO signups (id, email) VALUES (?, ?)"
	_, err = con.Exec(ctx, query, id, email)
	if err != nil {
		return "", err
	}
	return id, nil
}
