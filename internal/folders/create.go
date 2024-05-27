package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, rq *http.Request) {

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

	id, err := Insert(h.db, folders)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	folders.ID = id

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(folders)
}

func Insert(db *sql.DB, folders *Folders) (int64, error) {
	stmt := `insert into "folders" ("parent_id", "name", "modified_at") VALUES ($1, $2, $3)`

	result, err := db.Exec(stmt, folders.ParentID, folders.Name, folders.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
