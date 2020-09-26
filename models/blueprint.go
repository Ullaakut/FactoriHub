package models

import (
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
)

type Blueprint struct {
	// String, the name of the item that was saved ("blueprint" in vanilla).
	Item  string `json:"item" db:"-"`
	Label string `json:"label" db:"-"`
	// The map version of the map the blueprint was created in.
	Version VersionString `json:"version" db:"-"`

	Icons    Icons    `json:"icons" db:"-"`
	Entities Entities `json:"entities" db:"-"`

	// The following fields are not stored in the database at the moment.
	LabelColor Color      `json:"label_color" db:"-"`
	Tiles      []Tile     `json:"tiles" db:"-"`
	Schedules  []Schedule `json:"schedules" db:"-"`
}

func (b Blueprint) TotalEntities() map[string]uint {
	var total = make(map[string]uint)
	for _, e := range b.Entities {
		typ := e.Type
		switch typ {
		// No type specified.
		case "":
			typ = "item"

		// Legacy type naming.
		case "input", "output":
			typ = "item"
		}

		total[typ+"/"+e.Name] += 1
	}
	return total
}

// ValidateCreate decodes a blueprint's info from its raw string and verifies whether it is valid.
func (b Blueprint) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	fmt.Println("Before checking labels & entities amount")

	if len(b.Entities) == 0 {
		return validate.NewErrors(), errors.New("missing entities")
	}

	if len(b.Label) == 0 {
		return validate.NewErrors(), errors.New("missing blueprint name")
	}

	fmt.Println("Before creating entities")

	//for _, e := range b.Entities {
	//	verrs, err := tx.ValidateAndCreate(e)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint reference: %w", err)
	//	}
	//}

	fmt.Println("After creating entities")

	return validate.NewErrors(), nil
}

// ValidateUpdate decodes a blueprint's info from its raw string and verifies whether it is valid.
func (b Blueprint) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	if len(b.Entities) == 0 {
		return validate.NewErrors(), errors.New("missing entities")
	}

	if len(b.Label) == 0 {
		return validate.NewErrors(), errors.New("missing blueprint name")
	}

	for _, e := range b.Entities {
		verrs, err := tx.ValidateAndUpdate(e)
		if err != nil {
			return validate.NewErrors(), errors.WithStack(err)
		}
		if verrs.HasAny() {
			return verrs, fmt.Errorf("unable to create blueprint reference: %w", err)
		}
	}

	return validate.NewErrors(), nil
}
