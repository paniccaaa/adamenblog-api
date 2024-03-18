package post

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	res "github.com/paniccaaa/adamenblog-api/internal/lib/api/response"
	"github.com/paniccaaa/adamenblog-api/internal/storage/postgres"
)

// func HandleUpdatePost(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op = "handlers.post.HandleUpdatePost"

// 		log := log.With(
// 			slog.String("op", op),
// 			slog.String("request_id", middleware.GetReqID(r.Context())),
// 		)

// 		postID := chi.URLParam(r, "id")

// 		id, err := strconv.Atoi(postID)
// 		if err != nil {
// 			log.Error("failed to fetch post id")
// 		}

// 		post := &postgres.Post{}
// 		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
// 			log.Error("failed to decode request body")

// 			render.Status(r, http.StatusBadRequest)
// 			render.JSON(w, r, res.Error("failed to decode read body"))

// 			return
// 		}

// 		updatedPost, err := storage.UpdatePost(id, post)
// 		if err != nil {
// 			log.Error("failed to update post")

// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, res.Error("failed to update post"))

// 			return
// 		}

// 		render.JSON(w, r, updatedPost)
// 	}
// }

func HandleUpdatePost(log *slog.Logger, storage *postgres.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.post.HandleUpdatePost"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		postID := chi.URLParam(r, "id")

		id, err := strconv.Atoi(postID)
		if err != nil {
			log.Error("failed to fetch post id")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, res.Error("id could not found"))
			return
		}

		post := &postgres.Post{}
		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
			log.Error("failed to decode request body")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, res.Error("failed to decode request body"))
			return
		}

		updatedPost, err := storage.UpdatePost(id, post)
		if err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				log.Error("post does not exist")

				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, res.Error("id could not found"))
			} else {
				log.Error("failed to update post")

				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, res.Error("failed to update post"))
			}
			return
		}

		render.JSON(w, r, updatedPost)
	}
}
