package files

import (
	"database/sql"
	"time"
)

func Modifier(db *sql.DB, id int64, file *File) error {
	file.ModifiedAt = time.Now()
	stmt := `update "files" set "name"=$1, "modified"=$2, deleted=$3 where id=$4`

	_, err := db.Exec(stmt, file.Name, file.ModifiedAt, file.Deleted, id)

	return err

}
