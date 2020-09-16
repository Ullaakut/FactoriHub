package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type BlueprintBook struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	EntryID uuid.UUID `db:"entry_id"`
	Entry   *Entry    `json:"-" belongs_to:"entry"`

	// 	String, the name of the item that was saved ("blueprint-book" in vanilla).
	Item string `json:"item" db:"item"`

	// 	String, the name of the blueprint set by the user.
	Label string `json:"label" db:"name"`

	// The actual content of the blueprint book, array of objects containing an "index"
	// key and 0-based value and a "blueprint" key with a Blueprint object as the value.
	BlueprintReferences []BlueprintReference `json:"blueprints" has_many:"blueprint_references"`

	Version string `json:"version" db:"version"`

	// The following fields are not stored in the database at the moment.

	// Index of the currently selected blueprint, 0-based.
	ActiveIndex int   `json:"active_index"`
	LabelColor  Color `json:"label_color"`
}
