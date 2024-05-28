package folders

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetRootSubFolder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, "Documents", time.Now(), time.Now(), false).
		AddRow(3, nil, "Images", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where "parent_id" is null and "deleted"=false`)).
		WithArgs().
		WillReturnRows(rows)

	_, err = getRootSubFolder(db)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}
