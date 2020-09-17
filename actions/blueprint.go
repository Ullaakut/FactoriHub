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

	bp := models.Entry{}
	if err := tx.Eager().Find(&bp, c.Param("bpid")); err != nil {
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
