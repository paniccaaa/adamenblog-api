package post

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	res "github.com/paniccaaa/adamenblog-api/internal/lib/api/response"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"
)

type Request struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Response struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Image string `json:"image"`
}

func HandleGetPosts(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.handleGetPosts"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		posts, err := storage.GetPosts()
		if err != nil {
			log.Error("failed to get post")

			render.Status(r, http.StatusInternalServerError)

			render.JSON(w, r, res.Error(err.Error()))

			return
		}

		render.JSON(w, r, posts)

	}
}
