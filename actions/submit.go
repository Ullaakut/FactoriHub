package actions

import (
	"factorihub/models"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/gobuffalo/buffalo"
)

// SubmitHandler is a handler to submit a blueprint and redirect to the home page.
func SubmitHandler(c buffalo.Context) error {
	// Allocate an empty Entry
	entry := &models.Entry{
		Labels: make([]string, 0),
	}
	//user := c.Value("current_user").(*models.User)

	// Bind entry to the html form elements
	if err := c.Bind(entry); err != nil {
		return errors.WithStack(err)
	}

	//err := tx.Transaction(func(tx *pop.Connection) error {
	entry.AuthorID, _ = uuid.FromString("c756b686-f97d-11ea-adc1-0242ac120002")
	if err := createEntry(c, entry); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "New entry added successfully.")

	// and redirect to the index page
	return c.Redirect(302, "/")
}

func createEntry(c buffalo.Context, entry *models.Entry) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	verrs, err := tx.ValidateAndCreate(entry)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("entry", entry)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	if entry.BlueprintBook != nil {
		return createBlueprintBook(c, tx, entry.BlueprintBook)
	}

	return createBlueprint(c, tx, entry.Blueprint)
}

func createBlueprintBook(c buffalo.Context, tx *pop.Connection, bpb *models.BlueprintBook) error {
	verrs, err := tx.ValidateAndCreate(bpb)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("blueprint_book", bpb)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	for _, ref := range bpb.BlueprintReferences {
		if err := createBlueprintReference(c, tx, &ref); err != nil {
			return err
		}
	}

	return nil
}

func createBlueprintReference(c buffalo.Context, tx *pop.Connection, bpr *models.BlueprintReference) error {
	verrs, err := tx.ValidateAndCreate(bpr)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("blueprint_reference", bpr)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	return createBlueprint(c, tx, &bpr.Blueprint)
}

func createBlueprint(c buffalo.Context, tx *pop.Connection, bp *models.Blueprint) error {
	verrs, err := tx.ValidateAndCreate(bp)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("blueprint", bp)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	for _, entity := range bp.Entities {
		if err := createEntity(c, tx, &entity); err != nil {
			return err
		}
	}

	for _, icon := range bp.Icons {
		if err := createIcon(c, tx, &icon); err != nil {
			return err
		}
	}

	return nil
}

func createEntity(c buffalo.Context, tx *pop.Connection, entity *models.Entity) error {
	verrs, err := tx.ValidateAndCreate(entity)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("entity", entity)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	return nil
}

func createIcon(c buffalo.Context, tx *pop.Connection, icon *models.Icon) error {
	verrs, err := tx.ValidateAndCreate(icon)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("icon", icon)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	return createSignalID(c, tx, &icon.Signal)
}

func createSignalID(c buffalo.Context, tx *pop.Connection, signal *models.SignalID) error {
	verrs, err := tx.ValidateAndCreate(signal)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("signal", signal)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	return nil
}
