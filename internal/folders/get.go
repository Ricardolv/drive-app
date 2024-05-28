package folders

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) Get(rw http.ResponseWriter, rq *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	folders, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO GET contebn

	rw.Header().Set("Content-Type", "application/json")

}

func Get(db *sql.DB, id int64) (*Folders, error) {
	stmt := `select * from "folders" where id=$1`
	row := db.QueryRow(stmt)

	var folders Folders
	err := row.Scan(&folders.ID, &folders.ParentID, &folders.Name,
		&folders.CreatedAt, &folders.ModifiedAt, &folders.Deleted)
	if err != nil {
		return nil, err
	}

	return &folders, nil
}
