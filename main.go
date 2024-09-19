package main

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
	"go-werk-it/view/page"
)

func main() {
	app := echo.New()
	app.GET("/", ShowLandingPage)
	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}


func ShowLandingPage(c echo.Context) error {
	return render(c, page.Landing()) 
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}