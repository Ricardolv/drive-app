package users

import (
	"database/sql"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, rq *http.Request) {

}

func Insert(db *sql.DB, user *User) (int64, error) {
	stmt := `insert into "users" ("name", "login", "password", "modified_at") VALUES (:name, :login, :password, :modified_at) VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(stmt, user.Name, user.Login, user.Password, user.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
