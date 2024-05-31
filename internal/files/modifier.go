package files

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modifier(rw http.ResponseWriter, rq *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(rq.Body).Decode(&file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = file.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = Modifier(h.db, int64(id), file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(file)
}

func Modifier(db *sql.DB, id int64, file *File) error {
	file.ModifiedAt = time.Now()
	stmt := `update "files" set "name"=$1, "modified"=$2, deleted=$3 where id=$4`

	_, err := db.Exec(stmt, file.Name, file.ModifiedAt, file.Deleted, id)

	return err

}
