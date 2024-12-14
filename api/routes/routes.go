package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	email_handler "github.com/arinji2/garconia/emails/templates"
	"github.com/arinji2/garconia/logging"
	"github.com/arinji2/garconia/sqlite"
	"github.com/arinji2/garconia/verify"
	"github.com/go-chi/render"
)

func AddEmailRoute(w http.ResponseWriter, r *http.Request) {
	sync.OnceFunc(func() {
		fmt.Println("Registered /api/add endpoint.")
	})
	db, err := sqlite.NewConnection()
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute*1)
	defer cancel()
	var data AddEmailRequest
	err = parseRequestBody(r, &data)
	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}
	id, err := addEmail(ctx, data.Email, db)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: signups.email" {
			render.Status(r, 409)
			render.PlainText(w, r, "Email already exists")
		}
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}
	newUserLogger := logging.NewUserLogger(data.Email, id, r.RemoteAddr, r.UserAgent())
	newUserLogger.Send()
	userData := &sqlite.UserData{
		ID:    id,
		Email: data.Email,
	}
	err = verify.VerifyEmail(db, userData)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, id)
}

func VerifyEmailRoute(w http.ResponseWriter, r *http.Request) {
	db, err := sqlite.NewConnection()
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute*1)
	defer cancel()
	var data VerifyEmailRequest
	err = parseRequestBody(r, &data)
	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, err.Error())
		return
	}
	verified, err := verifyEmail(ctx, data.Code, data.Email, db)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	if verified {
		newVerifiedLogger := logging.NewVerificationLogger(data.Email, r.RemoteAddr)
		newVerifiedLogger.Send()
		email_handler.SendConfirmationEmail(data.Email)
		render.Status(r, 200)
	} else {
		newFailedVerificationLogger := logging.NewFailedVerificationLogger(data.Email, r.RemoteAddr)
		newFailedVerificationLogger.Send()
		render.Status(r, 400)
	}
}

// ParseRequestBody parses the JSON body of a request into the provided struct.
func parseRequestBody(r *http.Request, data interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(data); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("request body is empty")
		}
		return errors.New("invalid JSON format")
	}

	return nil
}
