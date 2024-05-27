package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modifier(rw http.ResponseWriter, rq *http.Request) {
	user := new(User)

	err := json.NewDecoder(rq.Body).Decode(user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Name == "" {
		http.Error(rw, ErrNameRequired.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(rq, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, int64(id), user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO GET id

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}

func Update(db *sql.DB, id int64, user *User) error {
	user.ModifiedAt = time.Now()

	stmt := `update "users" set "name"=$1, "modified"=$2 where id=$3`

	_, err := db.Exec(stmt, user.Name, user.ModifiedAt, id)

	return err
}
