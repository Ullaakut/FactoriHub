package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Icon is used by pop to map your icons database table to your go code.
type Icon struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	EntityID uuid.UUID `json:"-" db:"entity_id"`
	Entity   *Entity   `json:"-" belongs_to:"entity"`

	Signal SignalID `json:"signal" has_one:"signal"`
	Index  int      `json:"index"`
}
