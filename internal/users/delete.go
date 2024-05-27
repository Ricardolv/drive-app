package users

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, rq *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/")
}

func Delete(db *sql.DB, id int64) error {

	stmt := `update "users" set "modified"=$1, deleted=true where id=$2`

	_, err := db.Exec(stmt, time.Now(), id)

	return err
}
