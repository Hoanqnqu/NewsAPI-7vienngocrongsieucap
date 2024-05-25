package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AppRouter(dummyHandler *DummyHandler, userHandlers *UserHandlers, categoryHandlers *CategoryHandlers, newsHandlers *NewsHandlers) *chi.Mux {
	router := chi.NewRouter()
	router.Use(Logger)
	router.Use(AdminMiddleware)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Get("/dummy", dummyHandler.Dummy)
	router.Post("/login", userHandlers.Login)
	router.Get("/users", userHandlers.GetAll)
	router.Post("/users", userHandlers.Insert)
	router.Put("/users/{id}", userHandlers.Update)
	router.Get("/categories", categoryHandlers.GetAll)
	router.Post("/categories", categoryHandlers.Insert)
	router.Put("/categories/{id}", categoryHandlers.Update)
	router.Get("/news", newsHandlers.GetAll)
	router.Post("/news", newsHandlers.Insert)
	router.Put("/news/{id}", newsHandlers.Update)
	return router
}
