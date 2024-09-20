package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "<h1>Hello API Call!</h1>")
}

