package utils

import (
	"fmt"
	"os"
)

func GenerateVerificationURL(id, email string) string {
	baseURL := os.Getenv("FRONTEND_URL")
	return fmt.Sprintf("%s/confirm?code=%s&email=%s", baseURL, id, email)
}
