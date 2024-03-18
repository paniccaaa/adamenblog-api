package post

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	res "github.com/paniccaaa/adamenblog-api/internal/lib/api/response"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"
)

func HandleCreatePost(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.HandleCreatePost"

		post := &postgres.Post{}

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
			log.Error("failed to decode request body")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, res.Error("failed to decode read body"))

			return
		}

		createdPost, err := storage.CreatePost(post)
		if err != nil {
			log.Error("failed to create post")

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, res.Error("failed to create post"))

			return
		}

		render.JSON(w, r, createdPost)
	}
}
