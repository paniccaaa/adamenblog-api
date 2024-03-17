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

	router.Get("/posts", post.HandleGetPosts(log, storage))

	return router
}
