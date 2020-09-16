package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Signal struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Signal SignalID `json:"signal" has_one:"signal"`
	Count  int      `json:"count"`
}

type SignalID struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	SignalID uuid.UUID `json:"-" db:"signal_id"`
	Signal   *Signal   `json:"-" belongs_to:"signal"`

	IconID uuid.UUID `json:"-" db:"icon_id"`
	Icon   *Icon     `json:"-" belongs_to:"icon"`

	Type string `json:"type"`
	Name string `json:"name"`
}
