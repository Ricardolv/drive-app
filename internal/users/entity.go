package users

import "time"

func New(name, login, password string) (*User, error) {
	now := time.Now()

	user := User{Name: name, Login: login, CreatedAt: now, ModifiedAt: now}

	return &user, nil
}

type User struct {
	ID         int64
	Name       string
	Login      string
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted    bool
	LastLogin  time.Time
}
