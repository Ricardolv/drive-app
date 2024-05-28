package files

import (
	"database/sql"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, rq *http.Request) {

}

func Insert(db *sql.DB, file *File) (int64, error) {
	stmt := `insert into "files" ("folders_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2, $3, $4, $5, $6)`

	result, err := db.Exec(stmt, file.FolderID, file.OwnerID, file.Name, file.Type, file.Path, file.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
