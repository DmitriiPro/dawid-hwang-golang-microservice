package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	r.Use(middleware.Logger)
	tokenMaker := handler.TokenMaker

	r.Route("/products", func(r chi.Router) {
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Post("/", handler.CreateProduct)
		r.Get("/", handler.ListProducts)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)

			r.Group(func(r chi.Router) {
				r.Use(GetAdminMiddlewareFunc(tokenMaker))
				r.Patch("/", handler.UpdateProduct)
				r.Delete("/", handler.DeleteProduct)
			})

		})
	})

	r.Group(func(r chi.Router) {
		r.Use(GetAuthMiddlewareFunc(tokenMaker))
		r.Get("/myorder", handler.getOrder)

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", handler.CreateOrder)
			r.With(GetAdminMiddlewareFunc(tokenMaker)).Get("/", handler.ListOrders)

			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", handler.DeleteOrder)
			})
		})

	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)

		r.Post("/login", handler.loginUser)

		r.Group(func(r chi.Router) {
			r.Use(GetAdminMiddlewareFunc(tokenMaker))
			r.Get("/", handler.ListUsers)

			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", handler.DeleteUser)
			})

		})

		r.Group(func(r chi.Router) {
			r.Use(GetAuthMiddlewareFunc(tokenMaker))

			r.Patch("/", handler.UpdateUser)
			r.Post("/logout", handler.logoutUser)
		})

	})

	r.Group(func(r chi.Router) {
		r.Use(GetAuthMiddlewareFunc(tokenMaker))

		r.Route("/tokens", func(r chi.Router) {
			//* обновление токена доступа
			r.Post("/renew", handler.renewAccessToken)
			r.Post("/revoke", handler.revokeSession)
		})

	})

	return r
}

func Start(addr string) error {
	 log.Printf("Starting server at %s", addr)
	return http.ListenAndServe(addr, r)
}
