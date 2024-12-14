package logging

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	Payload WebhookPayload
}

var colors = map[string]int{
	"Green":  0x16a34a,
	"Brand":  0xa31621,
	"Purple": 0x800080,
}

func NewUserLogger(userEmail, userID, userIP, userAgent string) *Logger {
	return &Logger{
		Payload: WebhookPayload{
			Embeds: []Embed{
				{
					Title:       "New User Signup",
					Description: "New User Signed Up to Garconia Email",
					Color:       colors["Brand"],
					Fields: []Field{
						{
							Name:  "User Email",
							Value: userEmail,
						},
						{
							Name:   "User ID",
							Value:  userID,
							Inline: true,
						},
						{
							Name:   "User IP",
							Value:  userIP,
							Inline: true,
						},
						{
							Name:  "User Agent",
							Value: userAgent,
						},
					},
				},
			},
		},
	}
}

func NewFailedVerificationLogger(userEmail, userIP string) *Logger {
	return &Logger{
		Payload: WebhookPayload{
			Embeds: []Embed{
				{
					Title:       "New User Failed Verification",
					Description: "New User Verified their Verification",
					Color:       colors["Purple"],
					Fields: []Field{
						{
							Name:   "User IP",
							Value:  userIP,
							Inline: true,
						},
						{
							Name:  "User Email",
							Value: userEmail,
						},
					},
				},
			},
		},
	}
}

func NewVerificationLogger(userEmail string) *Logger {
	return &Logger{
		Payload: WebhookPayload{
			Embeds: []Embed{
				{
					Title:       "New User Verified Email",
					Description: "New User Verified their Email",
					Color:       colors["Green"],
					Fields: []Field{
						{
							Name:  "User Email",
							Value: userEmail,
						},
					},
				},
			},
		},
	}
}

func (l *Logger) Send() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	payloadBytes, err := json.Marshal(l.Payload)
	if err != nil {
		log.Println("Error marshalling payload ", err)
		return
	}
	webhookURL := os.Getenv("WEBHOOK_URL")
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("Error creating request: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()
}
