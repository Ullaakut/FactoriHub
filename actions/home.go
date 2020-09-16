package actions

import (
	"factorihub/models"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	// ?type=blueprints&version=1.0.0&label=mall&sort=latest
	//typeFilter := c.Param("type") // TODO: Unused for now.
	labelFilter := c.Param("label")
	versionFilter := c.Param("version")
	sort := c.Param("sort")
	query := c.Param("query")

	var blueprints []models.BlueprintData
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	q.Select("*").
		Where("labels LIKE '%' || ? || '%'", labelFilter).
		Where("version == ?", versionFilter).
		Where("name LIKE '%' || ? || '%'", query).
		Order(sort + " desc")
	if err := q.All(blueprints); err != nil {
		return errors.WithStack(err)
	}

	var versions models.Versions
	if err := q.All(versions); err != nil {
		return errors.WithStack(err)
	}

	var labels models.Labels
	if err := q.All(labels); err != nil {
		return errors.WithStack(err)
	}

	c.Set("versions", versions)
	c.Set("labels", labels)
	c.Set("blueprints", blueprints)

	return c.Render(http.StatusOK, r.HTML("index.html"))
}
