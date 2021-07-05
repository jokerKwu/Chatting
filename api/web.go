package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RedirectIndexPage(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/api/v1/index.html")
}
