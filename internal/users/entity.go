package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	ErrPasswordRequired = errors.New("password is required and can't be blank")

	ErrPasswordLen = errors.New("password must have at least 6 characters")

	ErrNameRequired = errors.New("name is required and can't be blank")

	ErrLoginRequired = errors.New("login is required and can't be blank")
)

func New(name, login, password string) (*User, error) {

	user := User{Name: name, Login: login, ModifiedAt: time.Now()}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"CreatedAt"`
	ModifiedAt time.Time `json:"ModifiedAt"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"LastLogin"`
}

func (u *User) SetPassword(password string) error {

	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) > 6 {
		return ErrPasswordLen
	}

	u.Password = fmt.Sprintf("%x", (md5.Sum([]byte(password))))

	return nil
}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameRequired
	}

	if u.Login == "" {
		return ErrLoginRequired
	}

	if u.Password == fmt.Sprintf("%x", (md5.Sum([]byte("")))) {
		return ErrPasswordRequired
	}

	return nil
}
