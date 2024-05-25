package rest

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AppRouter(dummyHandler *DummyHandler, userHandlers *UserHandlers, categoryHandlers *CategoryHandlers, newsHandlers *NewsHandlers) *chi.Mux {
	router := chi.NewRouter()
	router.Use(Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Post("/login", userHandlers.Login)

	router.Group(func(adminRouter chi.Router) {
		adminRouter.Use(AdminMiddleware)

		// User routes
		adminRouter.Get("/users", userHandlers.GetAll)
		adminRouter.Post("/users", userHandlers.Insert)
		adminRouter.Put("/users/{id}", userHandlers.Update)

		// Category routes
		adminRouter.Get("/categories", categoryHandlers.GetAll)
		adminRouter.Post("/categories", categoryHandlers.Insert)
		adminRouter.Put("/categories/{id}", categoryHandlers.Update)

		// News routes
		adminRouter.Get("/news", newsHandlers.GetAll)
		adminRouter.Post("/news", newsHandlers.Insert)
		adminRouter.Put("/news/{id}", newsHandlers.Update)
	})

	router.Group(func(userRouter chi.Router) {
		userRouter.Use(UserMiddleware)
		userRouter.Get("/dummy", dummyHandler.Dummy)
		userRouter.Post("/like/{newsId}", userHandlers.Like)
		userRouter.Post("/unlike/{newsId}", userHandlers.Unlike)
		userRouter.Post("/dislike/{newsId}", userHandlers.DisLike)
		userRouter.Post("/unDislike/{newsId}", userHandlers.UnDisLike)
	})

	return router
}
