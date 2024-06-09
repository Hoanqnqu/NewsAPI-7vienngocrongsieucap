package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func AppRouter(dummyHandler *DummyHandler, userHandlers *UserHandlers, categoryHandlers *CategoryHandlers, newsHandlers *NewsHandlers, s3Handler *UploadHandlers) *chi.Mux {
	router := chi.NewRouter()
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Post("/image/upload", s3Handler.Upload)
	router.Get("/news", newsHandlers.GetAll)
	router.Post("/login", userHandlers.Login)
	router.Get("/news/{newsId}", newsHandlers.GetNewsByID)
	router.Post("/adminlogin", userHandlers.AdminLogin)
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
		// adminRouter.Get("/news", newsHandlers.GetAll)
		adminRouter.Post("/news", newsHandlers.Insert)
		adminRouter.Put("/news/{id}", newsHandlers.Update)
	})

	router.Group(func(userRouter chi.Router) {
		userRouter.Use(UserMiddleware)
		userRouter.Get("/dummy", dummyHandler.Dummy)
		userRouter.Post("/like/{newsId}", userHandlers.Like)
		userRouter.Post("/dislike/{newsId}", userHandlers.Dislike)
		userRouter.Post("/save/{newsId}", userHandlers.Save)
		userRouter.Get("/saved", userHandlers.GetSavedNews)
	})

	return router
}
