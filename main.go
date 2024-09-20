package main

import (
	"database/sql"
	"go-werk-it/api"
	"go-werk-it/db"

	// "go-werk-it/routes"
	"go-werk-it/view/component"
	"go-werk-it/view/message"
	"go-werk-it/view/page"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)



func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close() // Ensure the DB is closed when the application exits

	app := echo.New()
	app.Static("/", "public")

	// routes.SetupAPIRoutes(app)

	app.GET("/", ShowLandingPage)

	app.GET("/test-string", HandleTestString)

	app.GET("/api/test", api.RegisterUser)
	  
	app.GET("/test-template", ShowTestTemplate)

	app.GET("/template/login", ShowLoginTemplate)
	app.GET("/template/register", ShowRegisterTemplate)

	app.GET("/login-template", ShowLoginTemplate)

	userHandler := UserHandler{DB: db}

	app.POST("/api/register", userHandler.HandleRegister)

	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) HandleRegister(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.String(http.StatusBadRequest, "Email and password cannot be empty")
	}

	// Insert new user into the database
	_, err := h.DB.Exec("INSERT INTO USER_AUTH (USERNAME, PASSWORD) VALUES (?, ?)", username, password)
	if err != nil {
		c.Logger().Errorf("Error inserting user into database: %v", err)
		return c.String(http.StatusInternalServerError, "Error registering user")
	}

	return c.String(http.StatusOK, "User registered successfully!")
}

func ShowLandingPage(c echo.Context) error {
	return render(c, page.Landing()) 
}

func ShowHomePage(c echo.Context) error {
	return render(c, page.Home()) 
}

func ShowLoginTemplate(c echo.Context) error {
	return render(c, component.Login())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func ShowRegisterTemplate(c echo.Context) error {
	return render(c, component.Register())
}

func HandleTestString(c echo.Context) error {
	return c.String(http.StatusOK, "<h1>String Test!</h1>")
}

func ShowTestTemplate(c echo.Context) error {
	return render(c, message.TestMessage()) 
}