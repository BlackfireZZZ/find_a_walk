package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            string `json:"error"`
	HTTPStatusCode int    `json:"-"`
}

type OKResponse struct {
	Status         string `json:"error"`
	HTTPStatusCode int    `json:"-"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}



func (e *OKResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error, statusCode int) render.Renderer {
	return &ErrResponse{
		Err:            err.Error(),
		HTTPStatusCode: statusCode,
	}
}
