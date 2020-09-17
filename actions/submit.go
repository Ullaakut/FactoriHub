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

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	//err := tx.Transaction(func(tx *pop.Connection) error {
	entry.AuthorID, _ = uuid.FromString("c756b686-f97d-11ea-adc1-0242ac120002")
	verrs, err := tx.Eager().ValidateAndCreate(entry)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("entry", entry)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("/upload"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "New entry added successfully.")

	// and redirect to the index page
	return c.Redirect(302, "/")
}
