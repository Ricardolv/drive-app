package folders

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Ricardolv/drive-app/internal/files"
	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, rq *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = deleteFiles(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO list folders

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

}

func Delete(db *sql.DB, id int64) error {
	stmt := `update "folders" set "modified"=$1, deleted=true where id=$2`

	_, err := db.Exec(stmt, time.Now(), id)

	return err
}

func deleteFiles(db *sql.DB, folderID int64) error {
	f, err := files.List(db, folderID)
	if err != nil {
		return err
	}

	removedFiles := make([]files.File, 0, len(f))
	for _, file := range f {
		file.Deleted = true
		err := files.Modifier(db, file.ID, &file)
		if err != nil {
			break
		}

		removedFiles = append(removedFiles, file)
	}

	if len(f) != len(removedFiles) {
		for _, file := range f {
			file.Deleted = true
			files.Modifier(db, file.ID, &file)
		}

		return err
	}

	return nil
}
