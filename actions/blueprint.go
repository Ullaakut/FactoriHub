package actions

import (
	"factorihub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"net/http"
	"reflect"
)

// BlueprintHandler is a default handler to serve up
// a blueprint page.
func BlueprintHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	entry := models.Entry{}
	if err := tx.Eager().Find(&entry, c.Param("bpid")); err != nil {
		return c.Error(404, err)
	}

	if err := entry.DecodeBlueprint(); err != nil {
		return c.Error(500, err)
	}
	c.Set("entry", entry)

	c.Logger().Infof("Entry: %+#v", entry)
	c.Logger().Infof("Blueprint.Blueprint: %+#v", entry.Blueprint)
	c.Logger().Infof("Blueprint.BlueprintBook: %+#v", entry.BlueprintBook)

	c.Set("notNil", func(c interface{}) bool {
		return !(c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()))
	})
	c.Set("versions", models.Versions{
		"1.0.0",
		"16.4.2",
	})
	c.Set("labels", []string{"tileable", "science", "mall", "rocket"})
	return c.Render(http.StatusOK, r.HTML("blueprint.html"))
}
