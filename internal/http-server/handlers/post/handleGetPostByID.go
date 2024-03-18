package post

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	res "github.com/paniccaaa/adamenblog-api/internal/lib/api/response"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"
)

func HandleGetPostByID(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.handleGetPostByID"
		postID := chi.URLParam(r, "id")

		id, err := strconv.Atoi(postID)
		if err != nil {
			log.Error("failed to fetch post id")
		}

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		post, err := storage.GetPostByID(id)
		if err != nil {
			log.Error("failed to get post")

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, res.Error("id could not found"))

			return
		}

		render.JSON(w, r, post)
	}
}
