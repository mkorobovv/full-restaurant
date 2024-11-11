package router

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	router *chi.Mux
}

func New() *Router {
	router := chi.NewRouter()

	r := Router{
		router: router,
	}

	return &r
}

func (r *Router) Router() *chi.Mux { return r.router }
