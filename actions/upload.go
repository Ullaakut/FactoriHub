package actions

import (
	"factorihub/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// UploadHandler is a handler to serve up
// the upload page.
func UploadHandler(c buffalo.Context) error {
	c.Set("entry", &models.Entry{
		Labels: make([]string, 0),
	})
	c.Set("labels", []string{"tileable", "science", "mall", "rocket"})

	return c.Render(http.StatusOK, r.HTML("upload.html"))
}
