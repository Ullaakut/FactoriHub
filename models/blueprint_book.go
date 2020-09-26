package models

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
)

type BlueprintBook struct {
	// 	String, the name of the blueprint set by the user.
	Label string `json:"label" db:"-"`

	// The actual content of the blueprint book, array of objects containing an "index"
	// key and 0-based value and a "blueprint" key with a Blueprint object as the value.
	BlueprintReferences BlueprintReferences `json:"blueprints" db:"-"`

	Version VersionString `json:"version" db:"-"`

	// The following fields are not stored in the database at the moment.

	// 	String, the name of the item that was saved ("blueprint-book" in vanilla).
	Item string `json:"item" db:"-"`

	// Index of the currently selected blueprint, 0-based.
	ActiveIndex int   `json:"active_index" db:"-"`
	LabelColor  Color `json:"label_color" db:"-"`
}

// ValidateCreate verifies whether a blueprint book and its contents are valid before creating them.
func (b BlueprintBook) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	if len(b.BlueprintReferences) == 0 {
		return validate.NewErrors(), errors.New("missing blueprint references")
	}

	//for _, bpr := range b.BlueprintReferences {
	//	verrs, err := tx.ValidateAndCreate(bpr)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint reference: %w", err)
	//	}
	//}

	return validate.NewErrors(), nil
}

// ValidateUpdate verifies whether a blueprint book and its contents are valid before updating them.
func (b BlueprintBook) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	if len(b.BlueprintReferences) == 0 {
		return validate.NewErrors(), errors.New("missing blueprint references")
	}

	//for _, bpr := range b.BlueprintReferences {
	//	verrs, err := tx.ValidateAndUpdate(bpr)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to update blueprint reference: %w", err)
	//	}
	//}

	return validate.NewErrors(), nil
}
