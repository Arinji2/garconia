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

	"github.com/arinji2/garconia/sqlite"
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
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, id)
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
