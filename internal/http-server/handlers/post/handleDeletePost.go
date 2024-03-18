package post

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"

	res "github.com/paniccaaa/adamenblog-api/internal/lib/api/response"
)

func HandleDeletePost(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.HandleDeletePost"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		postID := chi.URLParam(r, "id")

		id, err := strconv.Atoi(postID)
		if err != nil {
			log.Error("failed to fetch post id")
		}

		if err := storage.DeletePost(id); err != nil {
			log.Error("failed to delete post")

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, res.Error("id could not found"))

			return
		}

		render.NoContent(w, r)
	}
}
