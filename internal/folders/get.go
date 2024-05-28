package folders

import (
	"database/sql"
	"encoding/json"
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

	f, err := GetFolder(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	c, err := GetFolderContent(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fc := FolderContent{Folder: *f, Content: c}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)

}

func GetFolder(db *sql.DB, id int64) (*Folder, error) {
	stmt := `select * from "folders" where id=$1`
	row := db.QueryRow(stmt)

	var folders Folder
	err := row.Scan(&folders.ID, &folders.ParentID, &folders.Name,
		&folders.CreatedAt, &folders.ModifiedAt, &folders.Deleted)
	if err != nil {
		return nil, err
	}

	return &folders, nil
}

func GetFolderContent(db *sql.DB, folderID int64) ([]FolderResource, error) {
	subFolders, err := getSubFolder(db, folderID)
	if err != nil {
		return nil, err
	}

	fr := make([]FolderResource, 0, len(subFolders))
	for _, sf := range subFolders {
		resource := FolderResource{
			ID:         sf.ID,
			Name:       sf.Name,
			Type:       "directory",
			CreatedAt:  sf.CreatedAt,
			ModifiedAt: sf.ModifiedAt,
		}
		fr = append(fr, resource)
	}
	return fr, nil
}

func getSubFolder(db *sql.DB, parentID int64) ([]Folder, error) {
	stmt := `select * from "folders" where "parent_id"=$1 and "deleted"=false`
	rows, err := db.Query(stmt, parentID)
	if err != nil {
		return nil, err
	}

	f := make([]Folder, 0)
	for rows.Next() {
		var folder Folder
		err := rows.Scan(&folder.ID, &folder.ParentID, &folder.Name, &folder.CreatedAt,
			&folder.ModifiedAt, &folder.Deleted)
		if err != nil {
			continue
		}

		f = append(f, folder)
	}

	return f, nil
}
