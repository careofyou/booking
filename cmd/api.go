package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
}
