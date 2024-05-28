package files

import (
	"errors"
	"time"
)

var (
	ErrOwnerRequired = errors.New("owner is required and can't be blank")

	ErrTypeRequired = errors.New("type is required and can't be blank")

	ErrNameRequired = errors.New("name is required and can't be blank")

	ErrPathRequired = errors.New("path is required and can't be blank")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"ownerID"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {

	if f.OwnerID == 0 {
		return ErrOwnerRequired
	}

	if f.Name == "" {
		return ErrNameRequired
	}

	if f.Type == "" {
		return ErrTypeRequired
	}

	if f.Path == "" {
		return ErrPathRequired
	}

	return nil
}
