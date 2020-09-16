package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`

	Entries Entries `json:"-" has_many:"entries"`
}
