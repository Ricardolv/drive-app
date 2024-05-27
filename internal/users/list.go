package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) List(rw http.ResponseWriter, rq *http.Request) {

	users, err := SelectAll(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(users)
}

func SelectAll(db *sql.DB) ([]User, error) {
	stmt := `select * from "users" where deleted = false`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Login, &user.Password,
			&user.CreatedAt, &user.ModifiedAt, &user.Deleted, &user.LastLogin)

		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil

}
