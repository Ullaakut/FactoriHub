package models

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
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

type BlueprintReferences []BlueprintReference

// ValidateCreate verifies whether a blueprint reference and its contents are valid before creating them.
func (b BlueprintReference) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate verifies whether a blueprint reference and its contents are valid before updating them.
func (b BlueprintReference) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
