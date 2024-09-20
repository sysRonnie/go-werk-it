package routes

import (
	"go-werk-it/api"
	"github.com/labstack/echo/v4"
)

func SetupAPIRoutes(app *echo.Echo) {
	app.GET("api/test", api.RegisterUser)
}

