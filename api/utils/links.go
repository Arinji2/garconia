package utils

import (
	"fmt"
	"os"
)

func GenerateVerificationURL(id string) string {
	baseURL := os.Getenv("BASE_URL")
	return fmt.Sprintf("%s/api/verify?id=%s", baseURL, id)
}
