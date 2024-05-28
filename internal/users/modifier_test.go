package users

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestModifier(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`update "users" set "name"=$1, "modified"=$2 where id=$3`)).
		WithArgs("Richard", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, 1, &User{Name: "Richard"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
