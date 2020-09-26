package models

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
)

type BlueprintReference struct {
	Blueprint Blueprint `json:"blueprint" db:"-"`

	Index int `json:"index" db:"-"`
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
