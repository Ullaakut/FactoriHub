package actions

import (
	"net/http"
	"reflect"

	"factorihub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	// ?type=blueprints&version=1.0.0&label=mall&sort=latest
	//typeFilter := c.Param("type") // TODO: Unused for now.
	//labelFilter := c.Param("label") // TODO: Unused for now.
	//versionFilter := c.Param("version")
	//sort := c.Param("sort")
	//query := c.Param("query")

	var entries models.Entries
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.Eager().PaginateFromParams(c.Params())
	if err := q.All(&entries); err != nil {
		return errors.WithStack(err)
	}

	for idx := range entries {
		if err := entries[idx].DecodeBlueprint(); err != nil {
			return errors.WithStack(err)
		}
	}

	c.Set("notNil", func(c interface{}) bool {
		return !(c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()))
	})
	c.Set("versions", models.Versions{
		"1.0.0",
		"0.4.2",
	})
	c.Set("labels", []string{"tileable", "science", "mall", "rocket"})
	c.Set("entries", entries)

	return c.Render(http.StatusOK, r.HTML("home.html"))
}
