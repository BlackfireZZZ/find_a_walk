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
	Status         string `json:"status"`
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

func OKRequest(status string, statusCode int) render.Renderer {
	return &OKResponse{
		Status:         status,
		HTTPStatusCode: statusCode,
	}
}

func ErrInvalidRequest(err error, statusCode int) render.Renderer {
	return &ErrResponse{
		Err:            err.Error(),
		HTTPStatusCode: statusCode,
	}
}
