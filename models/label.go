package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Label is used by pop to map your labels database table to your go code.
type Label struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	EntryID uuid.UUID `db:"entry_id"`
	Entry   *Entry    `json:"-" belongs_to:"entry"`

	Value string `db:"value"`
}

// Labels represents multiple labels.
type Labels []Label
