package utils

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/arinji2/garconia/sqlite"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@#$%!"

// generates a 10 digit ID
func GenerateID(ctx context.Context, con *sqlite.Connection, table string) (string, error) {
	for {
		id := generateRandomString(10)
		exists, err := checkIDExists(ctx, con, id, table)
		if err != nil {
			return "", err
		}

		if !exists {
			return id, nil
		}
	}
}

func checkIDExists(ctx context.Context, con *sqlite.Connection, id string, table string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE id = ?", table)
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
