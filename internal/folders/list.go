package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Ricardolv/drive-app/internal/files"
)

func (h *handler) List(rw http.ResponseWriter, rq *http.Request) {

	c, err := GetRootFolderContent(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fc := FolderContent{
		Folder: Folder{
			Name: "root",
		},
		Content: c}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)

}

func GetRootFolderContent(db *sql.DB) ([]FolderResource, error) {
	subFolders, err := getRootSubFolder(db)
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

	folderList, err := files.ListRoot(db)
	if err != nil {
		return nil, err
	}

	for _, f := range folderList {
		resource := FolderResource{
			ID:         f.ID,
			Name:       f.Name,
			Type:       f.Type,
			CreatedAt:  f.CreatedAt,
			ModifiedAt: f.ModifiedAt,
		}
		fr = append(fr, resource)
	}

	return fr, nil
}

func getRootSubFolder(db *sql.DB) ([]Folder, error) {
	stmt := `select * from "folders" where "parent_id" is null and "deleted"=false`
	rows, err := db.Query(stmt)
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
