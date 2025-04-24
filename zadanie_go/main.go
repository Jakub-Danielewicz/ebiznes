package main

import(
	"zadanie_go/db"
	"zadanie_go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.ConnectDB()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
