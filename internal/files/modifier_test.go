package files

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func TestModifierHttp(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db, nil, nil}

	f := File{
		ID:         1,
		Name:       "Name file png",
		OwnerID:    1,
		Type:       "images/png",
		Path:       "/",
		ModifiedAt: time.Now(),
		Deleted:    false,
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, f.OwnerID, f.Name, f.Type, f.Path, time.Now(), f.ModifiedAt, f.Deleted)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where id=$1`)).
		WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(`update "files" set "name"=$1, "modified"=$2, deleted=$3 where id=$4`)).
		WithArgs("Name file png", AnyTime{}, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Modifier(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}

func TestModifierDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`update "files" set "name"=$1, "modified"=$2, deleted=$3 where id=$4`)).
		WithArgs("Name file png", AnyTime{}, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Modifier(db, 1, &File{Name: "Name file png"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
