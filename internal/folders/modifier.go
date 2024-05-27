package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modifier(rw http.ResponseWriter, rq *http.Request) {
	folders := new(Folders)

	err := json.NewDecoder(rq.Body).Decode(folders)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = folders.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, int64(id), folders)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO GET id

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(folders)
}

func Update(db *sql.DB, id int64, folders *Folders) error {
	folders.ModifiedAt = time.Now()
	stmt := `update "folders" set "name"=$1, "modified"=$2 where id=$3`

	_, err := db.Exec(stmt, folders.Name, folders.ModifiedAt, id)

	return err
}
