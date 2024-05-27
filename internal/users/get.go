package users

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

	user, err := Get(h.db, int64(id))
	if err != nil {
		//TODO exists user ?
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}

func Get(db *sql.DB, id int64) (*User, error) {

	stmt := `select * from "users" where id=$1`
	row := db.QueryRow(stmt)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password,
		&user.CreatedAt, &user.ModifiedAt, &user.Deleted, &user.LastLogin)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
