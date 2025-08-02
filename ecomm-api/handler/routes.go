package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.ListProducts)
		r.Post("/", handler.CreateProduct)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)
			r.Patch("/", handler.UpdateProduct)
			r.Delete("/", handler.DeleteProduct)
		})
	})

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.CreateOrder)
		r.Get("/", handler.ListOrders)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getOrder)
			r.Delete("/", handler.DeleteOrder)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)
		r.Get("/", handler.ListUsers)
		r.Patch("/", handler.UpdateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", handler.DeleteUser)
		})

		r.Route("/login", func(r chi.Router) {
			r.Post("/", handler.loginUser)
		})

		r.Route("/logout", func(r chi.Router) {
			r.Post("/", handler.logoutUser)
		})

	})

	r.Route("/tokens", func(r chi.Router) {
		//* обновление токена доступа
		r.Route("/renew", func(r chi.Router) {
			r.Post("/", handler.renewAccessToken)
		})

		r.Route("/revoke/{id}", func(r chi.Router) {
			r.Post("/", handler.revokeSession)
		})
	})

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
