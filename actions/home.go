package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

// HomeAbout default implementation.
func HomeAbout(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/about.html"))
}

// HomeContact default implementation.
func HomeContact(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/contact.html"))
}

