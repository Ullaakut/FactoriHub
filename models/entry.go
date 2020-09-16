package models

import (
	"errors"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
)

// Entry is used by pop to map your entries database table to your go code.
type Entry struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	AuthorID uuid.UUID `json:"author_id" db:"user_id"`
	Author   User      `json:"-" belongs_to:"user"`

	Title         string         `json:"-" db:"title"`
	Description   string         `json:"-" db:"description"`
	Favorites     uint           `json:"-" db:"favorites"`
	Image         string         `json:"-" db:"image_path"`
	RawString     string         `json:"-" db:"raw_string"`
	RedditThread  string         `json:"-" db:"reddit_thread"`
	Labels        Labels         `json:"-" has_many:"labels"`
	BlueprintBook *BlueprintBook `json:"blueprints" has_one:"blueprint_book"`
	Blueprint     *Blueprint     `json:"blueprint" has_one:"blueprint"`
}

// Entries is not required by pop and may be deleted
type Entries []Entry

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
func (e *Entry) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	if e.BlueprintBook == nil && e.Blueprint == nil {
		return validate.NewErrors(), errors.New("missing blueprint")
	}
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Entry) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	if e.BlueprintBook == nil && e.Blueprint == nil {
		return validate.NewErrors(), errors.New("missing blueprint")
	}
	return validate.NewErrors(), nil
}
