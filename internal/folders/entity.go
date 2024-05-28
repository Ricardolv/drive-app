package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequired = errors.New("name is required and can't be blank")
)

func New(name string, parentID int64) (*Folder, error) {
	folders := Folder{
		Name:       name,
		ParentID:   parentID,
		ModifiedAt: time.Now(),
	}

	err := folders.Validate()
	if err != nil {
		return nil, err
	}

	return &folders, nil
}

type Folder struct {
	ID         int64     `json:"id"`
	ParentID   int64     `json:"parentID"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Deleted    bool      `json:"-"`
}

func (folders *Folder) Validate() error {

	if folders.Name == "" {
		return ErrNameRequired
	}

	return nil
}

type FolderContent struct {
	Folder  Folder           `json:"folder"`
	Content []FolderResource `json:"content"`
}

type FolderResource struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
