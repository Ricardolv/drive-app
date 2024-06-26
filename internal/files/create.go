package files

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ricardolv/drive-app/internal/queue"
)

func (h *handler) Create(rw http.ResponseWriter, rq *http.Request) {
	rq.ParseMultipartForm(32 << 20)

	file, fileHeader, err := rq.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	path := fmt.Sprintf("/%s", fileHeader.Filename)

	err = h.bucket.Upload(file, path)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := New(1, fileHeader.Filename, fileHeader.Header.Get("content-type"), path)
	if err != nil {
		h.bucket.Delete(path)

		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	folderID := rq.Form.Get("folderID")
	if folderID == "" {
		fid, err := strconv.Atoi(folderID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		entity.FolderID = int64(fid)
	}

	id, err := Insert(h.db, entity)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	entity.ID = id

	message := queue.Message{
		Filename: fileHeader.Filename,
		Path:     path,
		ID:       int(id),
	}

	msg, err := message.Marshal()
	if err != nil {
		// TODO: rollback

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.queue.Publish(msg)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(entity)
}

func Insert(db *sql.DB, file *File) (int64, error) {
	stmt := `insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2, $3, $4, $5, $6)`

	result, err := db.Exec(stmt, file.FolderID, file.OwnerID, file.Name, file.Type, file.Path, file.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
