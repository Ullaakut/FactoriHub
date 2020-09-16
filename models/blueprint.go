package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Blueprint struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// String, the name of the item that was saved ("blueprint" in vanilla).
	Item  string `json:"item" db:"item"`
	Label string `json:"label" db:"name"`
	// The map version of the map the blueprint was created in.
	Version VersionString `json:"version" db:"version"`

	Icons    []Icon   `json:"icons" has_many:"icons"`
	Entities []Entity `json:"entities" has_many:"entities"`

	EntryID uuid.UUID `db:"entry_id"`
	Entry   *Entry    `json:"-" belongs_to:"entry"`

	BlueprintReferenceID uuid.UUID           `db:"blueprint_reference_id"`
	BlueprintReference   *BlueprintReference `json:"-" belongs_to:"blueprint_reference"`

	// The following fields are not stored in the database at the moment.
	LabelColor Color      `json:"label_color"`
	Tiles      []Tile     `json:"tiles"`
	Schedules  []Schedule `json:"schedules"`
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
