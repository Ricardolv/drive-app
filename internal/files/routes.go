package files

import (
	"database/sql"

	"github.com/Ricardolv/drive-app/internal/bucket"
	"github.com/Ricardolv/drive-app/internal/queue"
	"github.com/go-chi/chi"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func SetRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Post("/", h.Create)
	r.Put("/{id}", h.Modifier)
	r.Delete("/{id}", h.Delete)

}
