package files

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	f, err := New(1, "Name file png", "png", "/")
	if err != nil {
		t.Error(err)
	}

	f.FolderID = 1

	mock.ExpectExec(regexp.QuoteMeta(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(1, 1, "Name file png", "png", "/", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
