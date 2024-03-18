package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/paniccaaa/adamenblog-api/internal/http-server/handlers/post"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"
)

func InitRouter(log *slog.Logger, storage *postgres.PostgresStore) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Route("/posts", func(r chi.Router) {
		r.Get("/", post.HandleGetPosts(log, storage))
		r.Get("/{id}", post.HandleGetPostByID(log, storage))

		r.Post("/", post.HandleCreatePost(log, storage))
		r.Patch("/{id}", post.HandleUpdatePost(log, storage))

		r.Delete("/{id}", post.HandleDeletePost(log, storage))
	})

	return router
}
