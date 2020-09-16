package actions

import (
	"factorihub/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("versions", models.Versions{
		{1, 0, 0, 0},
		{0, 16, 4, 2},
	})
	c.Set("labels", []string{
		"tileable",
		"mall",
		"science",
		"balancer",
		"rocket",
		"power",
		"furnace",
		"defense",
	})

	return c.Render(http.StatusOK, r.HTML("index.html"))
}
