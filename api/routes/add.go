package routes

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/arinji2/garconia/sqlite"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@#$%!"

type AddEmailRequest struct {
	Email string `json:"email"`
}

func addEmail(ctx context.Context, email string, con *sqlite.Connection) (string, error) {
	id, err := GenerateID(ctx, con)
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	query := "INSERT INTO signups (id, email) VALUES (?, ?)"
	_, err = con.Exec(ctx, query, id, email)
	if err != nil {
		return "", err
	}
	return id, nil
}

// generates a 10 digit ID
func GenerateID(ctx context.Context, con *sqlite.Connection) (string, error) {
	for {
		id := generateRandomString(10)
		exists, err := checkIDExists(ctx, con, id)
		if err != nil {
			return "", err
		}

		if !exists {
			return id, nil
		}
	}
}

func checkIDExists(ctx context.Context, con *sqlite.Connection, id string) (bool, error) {
	query := "SELECT COUNT(1) FROM signups WHERE id = ?"
	rows, err := con.Query(ctx, query, id)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
