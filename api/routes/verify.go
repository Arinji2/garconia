package routes

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/arinji2/garconia/sqlite"
)

type VerifyEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func verifyEmail(ctx context.Context, code, email string, con *sqlite.Connection) (bool, error) {
	query := `SELECT 
            v.id as verification_id,
            v.user_id,
            s.email as user_email
            FROM 
            verifications v
            JOIN 
            signups s 
            ON
            v.user_id = s.id
            WHERE
            v.id = ?
            `
	var verificationID string
	var userID string
	var userEmail string
	rows, err := con.Query(ctx, query, code)
	if err != nil {
		fmt.Println("Query error:", err)
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, errors.New("no verification found")
	}

	err = rows.Scan(&verificationID, &userID, &userEmail)
	if err != nil {
		fmt.Println("Scan error:", err)
		return false, err
	}
	fmt.Println(verificationID, userEmail)
	if verificationID == "" || userEmail == "" {
		return false, errors.New("verification not found")
	}

	if verificationID == code && userEmail == email {
		err = cleanupVerifications(ctx, verificationID, userID, con)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func cleanupVerifications(ctx context.Context, verificationID, userID string, con *sqlite.Connection) error {
	verificationQuery := `DELETE FROM verifications WHERE id = ?`
	userQuery := `UPDATE signups SET verified = 1 WHERE id = ?`
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := con.Exec(ctx, verificationQuery, verificationID)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := con.Exec(ctx, userQuery, userID)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Wait()
	close(errorChan)
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}
