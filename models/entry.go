package models

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"time"
)

// Entry represents an entry which can contain either a BP or a BP book.
type Entry struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	AuthorID uuid.UUID `json:"author_id" db:"user_id"`
	Author   User      `json:"-" belongs_to:"user"`

	Title         string         `form:"title" json:"-" db:"title"`
	Description   string         `form:"description" json:"-" db:"description"`
	Favorites     uint           `json:"-" db:"favorites"`
	Image         string         `json:"-" db:"image"`
	RawString     string         `form:"raw_string" json:"-" db:"raw_string"`
	RedditThread  string         `form:"reddit_thread" json:"-" db:"reddit_thread"`
	Labels        slices.String  `form:"labels" json:"-" db:"labels"`
	Version       string         `json:"-" db:"version"`
	BlueprintBook *BlueprintBook `json:"blueprint_book" has_one:"blueprint_book"`
	Blueprint     *Blueprint     `json:"blueprint" has_one:"blueprint"`
}

// Entries is not required by pop and may be deleted
type Entries []Entry

func (e *Entry) VersionString() VersionString {
	return VersionString(e.Version)
}

// DecodeBlueprint decodes an entry's string into a blueprint/blueprint book.
// Skip the first byte, base64 decode the string, and finally decompress using zlib inflate.
func (e *Entry) DecodeBlueprint() error {
	if len(e.RawString) < 2 {
		return fmt.Errorf("malformed blueprint string of length %d", len(e.RawString))
	}

	data, err := base64.StdEncoding.DecodeString(e.RawString[1:])
	if err != nil {
		return err
	}

	binData := bytes.NewReader(data)
	r, err := zlib.NewReader(binData)
	if err != nil {
		return err
	}

	jsonData := &bytes.Buffer{}
	io.Copy(jsonData, r)
	defer r.Close()

	_ = ioutil.WriteFile("/tmp/"+e.Title+".json", jsonData.Bytes(), 0655)

	if err := json.Unmarshal(jsonData.Bytes(), &e); err != nil {
		return err
	}

	if e.BlueprintBook != nil {
		e.Version = string(e.BlueprintBook.Version)
	}

	if e.Blueprint != nil {
		e.Version = string(e.Blueprint.Version)
	}

	return nil
}

// ValidateCreate decodes an entry's info from its raw string and verifies whether it is valid.
func (e *Entry) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	if err := e.DecodeBlueprint(); err != nil {
		return validate.NewErrors(), fmt.Errorf("invalid blueprint string: %w", err)
	}

	fmt.Printf("DECODED ENTRY: %v\n", e)
	if e.BlueprintBook != nil {
		for _, bp := range e.BlueprintBook.BlueprintReferences {
			fmt.Printf("DECODED BOOK: %v\n", bp)
		}
	} else if e.Blueprint != nil {
		fmt.Printf("DECODED BOOK: %v\n", e.Blueprint)
	}

	if e.BlueprintBook == nil && e.Blueprint == nil {
		return validate.NewErrors(), errors.New("missing blueprint")
	}

	fmt.Println("Before creating BP")

	//if e.BlueprintBook.Version != "" {
	//	verrs, err := tx.ValidateAndCreate(e.BlueprintBook)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint book: %w", err)
	//	}
	//} else {
	//	verrs, err := tx.ValidateAndCreate(e.Blueprint)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint: %w", err)
	//	}
	//}

	fmt.Println("After creating BP")

	return validate.NewErrors(), nil
}

// ValidateUpdate decodes an entry's info from its raw string and verifies whether it is valid.
func (e *Entry) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	if err := e.DecodeBlueprint(); err != nil {
		return validate.NewErrors(), fmt.Errorf("invalid blueprint string: %w", err)
	}

	if e.BlueprintBook == nil && e.Blueprint == nil {
		return validate.NewErrors(), errors.New("missing blueprint")
	}

	//if e.BlueprintBook.Version != "" {
	//	verrs, err := tx.ValidateAndUpdate(e.BlueprintBook)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint book: %w", err)
	//	}
	//} else {
	//	verrs, err := tx.ValidateAndUpdate(e.Blueprint)
	//	if err != nil {
	//		return validate.NewErrors(), errors.WithStack(err)
	//	}
	//	if verrs.HasAny() {
	//		return verrs, fmt.Errorf("unable to create blueprint: %w", err)
	//	}
	//}

	return validate.NewErrors(), nil
}
