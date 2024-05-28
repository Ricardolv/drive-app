package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, rq *http.Request) {

	user := new(User)

	err := json.NewDecoder(rq.Body).Decode(user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	user.SetPassword(user.Password)

	err = user.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := Insert(h.db, user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}

func Insert(db *sql.DB, user *User) (int64, error) {
	stmt := `insert into "users" ("name", "login", "password", "modified_at") VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(stmt, user.Name, user.Login, user.Password, user.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
