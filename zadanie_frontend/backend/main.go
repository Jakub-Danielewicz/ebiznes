package main

import(
	"backend/db"
	"backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	db.ConnectDB()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
