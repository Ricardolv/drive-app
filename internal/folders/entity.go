package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequired = errors.New("name is required and can't be blank")
)

func New(name string, parentID int64) (*Folders, error) {
	folders := Folders{
		Name:     name,
		ParentID: parentID,
	}

	err := folders.Validate()
	if err != nil {
		return nil, err
	}

	return &folders, nil
}

type Folders struct {
	ID         int64     `json:"id"`
	ParentID   int64     `json:"ParentID"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"CreatedAt"`
	ModifiedAt time.Time `json:"ModifiedAt"`
	Deleted    bool      `json:"-"`
}

func (folders *Folders) Validate() error {

	if folders.Name == "" {
		return ErrNameRequired
	}

	return nil
}
