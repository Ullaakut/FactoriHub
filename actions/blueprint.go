package actions

import (
	"factorihub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"net/http"
)

// BlueprintHandler is a default handler to serve up
// a blueprint page.
func BlueprintHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	bp, err := getEntry(c, tx)
	if err != nil {
		return c.Error(404, err)
	}

	c.Logger().Infof("Blueprint: %+#v", bp)
	c.Logger().Infof("Blueprint.Blueprint: %+#v", bp.Blueprint)
	c.Logger().Infof("Blueprint.BlueprintBook: %+#v", bp.BlueprintBook)
	c.Set("blueprint", bp)

	c.Set("versions", models.Versions{
		"1.0.0",
		"16.4.2",
	})
	c.Set("labels", []string{"tileable", "science", "mall", "rocket"})
	return c.Render(http.StatusOK, r.HTML("blueprint.html"))
}

func getEntry(c buffalo.Context, tx *pop.Connection) (models.Entry, error) {
	entry := models.Entry{}
	if err := tx.Eager().Find(&entry, c.Param("bpid")); err != nil {
		return entry, err
	}

	var err error
	if entry.BlueprintBook != nil {
		entry.BlueprintBook, err = getBlueprintBook(tx, entry.BlueprintBook)
		if err != nil {
			return entry, err
		}
	}

	entry.Blueprint, err = getBlueprint(tx, entry.Blueprint)
	return entry, err
}

func getBlueprintBook(tx *pop.Connection, bpb *models.BlueprintBook) (*models.BlueprintBook, error) {
	if err := tx.Eager().Find(&(*bpb), bpb.ID); err != nil {
		return bpb, err
	}

	return bpb, nil
}

func getBlueprint(tx *pop.Connection, bp *models.Blueprint) (*models.Blueprint, error) {
	if err := tx.Eager().Find(&(*bp), bp.ID); err != nil {
		return bp, err
	}

	return bp, nil
}
