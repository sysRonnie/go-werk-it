package routes

import (
	"go-werk-it/view/component"
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
)

func SetupTemplateRoutes(app *echo.Echo) {

	app.GET("/template/login", ShowLoginTemplate)
	app.GET("/template/register", ShowRegisterTemplate)
}


func ShowLoginTemplate(c echo.Context) error {
	return render(c, component.Login())
}

func ShowRegisterTemplate(c echo.Context) error {
	return render(c, component.Register())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
