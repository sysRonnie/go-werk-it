package main

import (
	"go-werk-it/view/page"
	"go-werk-it/view/message"
	"go-werk-it/view/component"
	"log"
	"net/http"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "public")
	app.GET("/", ShowLandingPage)

	app.GET("/test-string", HandleTestString)
	app.GET("/test-template", ShowTestTemplate)

	app.GET("/login-template", ShowLoginTemplate)

	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}


func ShowLandingPage(c echo.Context) error {
	return render(c, page.Landing()) 
}

func ShowLoginTemplate(c echo.Context) error {
	return render(c, component.Login())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}


func HandleTestString(c echo.Context) error {
	return c.String(http.StatusOK, "<h1>String Test!</h1>")
}

func ShowTestTemplate(c echo.Context) error {
	return render(c, message.TestMessage()) 
}