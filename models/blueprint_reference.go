package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type BlueprintReference struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	BlueprintBookID uuid.UUID      `db:"blueprint_book_id"`
	BlueprintBook   *BlueprintBook `belongs_to:"blueprint_book"`

	Blueprint Blueprint `json:"blueprint" has_one:"blueprint"`

	Index int `json:"index"`
}
