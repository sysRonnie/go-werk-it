package routes

import (
	
	"go-werk-it/view/page"
	"github.com/labstack/echo/v4"
	
)

func SetupPageRoutes(app *echo.Echo) {
	app.Static("/", "public")
	
	app.GET("/", ShowLandingPage)
	
}

func ShowLandingPage(c echo.Context) error {
	return render(c, page.Landing())
}

